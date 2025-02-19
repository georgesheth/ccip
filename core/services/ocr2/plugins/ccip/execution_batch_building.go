package ccip

import (
	"context"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merklemulti"
)

func getProofData(
	ctx context.Context,
	sourceReader ccipdata.OnRampReader,
	interval ccipdata.CommitStoreInterval,
) (sendReqsInRoot []ccipdata.Event[internal.EVM2EVMMessage], leaves [][32]byte, tree *merklemulti.Tree[[32]byte], err error) {
	sendReqs, err := sourceReader.GetSendRequestsBetweenSeqNums(
		ctx,
		interval.Min,
		interval.Max,
		0, // no need for confirmations, commitReport was already confirmed and we need all msgs in it
	)
	if err != nil {
		return nil, nil, nil, err
	}
	leaves = make([][32]byte, 0, len(sendReqs))
	for _, req := range sendReqs {
		leaves = append(leaves, req.Data.Hash)
	}
	tree, err = merklemulti.NewTree(hashlib.NewKeccakCtx(), leaves)
	if err != nil {
		return nil, nil, nil, err
	}
	return sendReqs, leaves, tree, nil
}

func buildExecutionReportForMessages(
	msgsInRoot []ccipdata.Event[internal.EVM2EVMMessage],
	leaves [][32]byte,
	tree *merklemulti.Tree[[32]byte],
	commitInterval ccipdata.CommitStoreInterval,
	observedMessages []ObservedMessage,
) (ccipdata.ExecReport, error) {
	innerIdxs := make([]int, 0, len(observedMessages))
	var messages []internal.EVM2EVMMessage
	var offchainTokenData [][][]byte
	for _, observedMessage := range observedMessages {
		if observedMessage.SeqNr < commitInterval.Min || observedMessage.SeqNr > commitInterval.Max {
			// We only return messages from a single root (the root of the first message).
			continue
		}
		innerIdx := int(observedMessage.SeqNr - commitInterval.Min)
		messages = append(messages, msgsInRoot[innerIdx].Data)
		offchainTokenData = append(offchainTokenData, observedMessage.TokenData)
		innerIdxs = append(innerIdxs, innerIdx)
	}

	merkleProof, err := tree.Prove(innerIdxs)
	if err != nil {
		return ccipdata.ExecReport{}, err
	}

	// any capped proof will have length <= this one, so we reuse it to avoid proving inside loop, and update later if changed
	return ccipdata.ExecReport{
		Messages:          messages,
		Proofs:            merkleProof.Hashes,
		ProofFlagBits:     abihelpers.ProofFlagsToBits(merkleProof.SourceFlags),
		OffchainTokenData: offchainTokenData,
	}, nil
}

// Validates the given message observations do not exceed the committed sequence numbers
// in the commitStoreReader.
func validateSeqNumbers(serviceCtx context.Context, commitStore ccipdata.CommitStoreReader, observedMessages []ObservedMessage) error {
	nextMin, err := commitStore.GetExpectedNextSequenceNumber(serviceCtx)
	if err != nil {
		return err
	}
	// observedMessages are always sorted by SeqNr and never empty, so it's safe to take last element
	maxSeqNumInBatch := observedMessages[len(observedMessages)-1].SeqNr

	if maxSeqNumInBatch >= nextMin {
		return errors.Errorf("Cannot execute uncommitted seq num. nextMin %v, seqNums %v", nextMin, observedMessages)
	}
	return nil
}

// Gets the commit report from the saved logs for a given sequence number.
func getCommitReportForSeqNum(ctx context.Context, commitStoreReader ccipdata.CommitStoreReader, seqNum uint64) (ccipdata.CommitStoreReport, error) {
	acceptedReports, err := commitStoreReader.GetAcceptedCommitReportsGteSeqNum(ctx, seqNum, 0)
	if err != nil {
		return ccipdata.CommitStoreReport{}, err
	}

	for _, acceptedReport := range acceptedReports {
		reportInterval := acceptedReport.Data.Interval
		if reportInterval.Min <= seqNum && seqNum <= reportInterval.Max {
			return acceptedReport.Data, nil
		}
	}

	return ccipdata.CommitStoreReport{}, errors.Errorf("seq number not committed")
}
