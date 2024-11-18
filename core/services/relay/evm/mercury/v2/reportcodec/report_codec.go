package reportcodec

import (
	"errors"
	"fmt"
	"math/big"

	pkgerrors "github.com/pkg/errors"

	ocrtypes "github.com/goplugin/plugin-libocr/offchainreporting2plus/types"

	"github.com/goplugin/plugin-common/pkg/logger"
	v2 "github.com/goplugin/plugin-common/pkg/types/mercury/v2"

	"github.com/goplugin/pluginv3.0/v2/core/services/relay/evm/mercury/utils"
	reporttypes "github.com/goplugin/pluginv3.0/v2/core/services/relay/evm/mercury/v2/types"
)

var ReportTypes = reporttypes.GetSchema()
var maxReportLength = 32 * len(ReportTypes) // each arg is 256 bit EVM word
var zero = big.NewInt(0)

var _ v2.ReportCodec = &ReportCodec{}

type ReportCodec struct {
	logger logger.Logger
	feedID utils.FeedID
}

func NewReportCodec(feedID [32]byte, lggr logger.Logger) *ReportCodec {
	return &ReportCodec{lggr, feedID}
}

func (r *ReportCodec) BuildReport(rf v2.ReportFields) (ocrtypes.Report, error) {
	var merr error
	if rf.BenchmarkPrice == nil {
		merr = errors.Join(merr, errors.New("benchmarkPrice may not be nil"))
	}
	if rf.LinkFee == nil {
		merr = errors.Join(merr, errors.New("linkFee may not be nil"))
	} else if rf.LinkFee.Cmp(zero) < 0 {
		merr = errors.Join(merr, fmt.Errorf("linkFee may not be negative (got: %s)", rf.LinkFee))
	}
	if rf.NativeFee == nil {
		merr = errors.Join(merr, errors.New("nativeFee may not be nil"))
	} else if rf.NativeFee.Cmp(zero) < 0 {
		merr = errors.Join(merr, fmt.Errorf("nativeFee may not be negative (got: %s)", rf.NativeFee))
	}
	if merr != nil {
		return nil, merr
	}
	reportBytes, err := ReportTypes.Pack(r.feedID, rf.ValidFromTimestamp, rf.Timestamp, rf.NativeFee, rf.LinkFee, rf.ExpiresAt, rf.BenchmarkPrice)
	return ocrtypes.Report(reportBytes), pkgerrors.Wrap(err, "failed to pack report blob")
}

func (r *ReportCodec) MaxReportLength(n int) (int, error) {
	return maxReportLength, nil
}

func (r *ReportCodec) ObservationTimestampFromReport(report ocrtypes.Report) (uint32, error) {
	decoded, err := r.Decode(report)
	if err != nil {
		return 0, err
	}
	return decoded.ObservationsTimestamp, nil
}

func (r *ReportCodec) Decode(report ocrtypes.Report) (*reporttypes.Report, error) {
	return reporttypes.Decode(report)
}

func (r *ReportCodec) BenchmarkPriceFromReport(report ocrtypes.Report) (*big.Int, error) {
	decoded, err := r.Decode(report)
	if err != nil {
		return nil, err
	}
	return decoded.BenchmarkPrice, nil
}