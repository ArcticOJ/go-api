package arctic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ArcticOJ/go-api-bindings/v0/types"
	"github.com/ArcticOJ/go-api-bindings/v0/types/common"
	"github.com/ArcticOJ/go-api-bindings/v0/types/submission"
	"github.com/mitchellh/mapstructure"
	"io"
)

func (c *Client) submitGeneric(ctx context.Context, problem, src, runtime string, streamed bool) (interface{}, error) {
	var submissionId uint32
	req := c.c.R().
		SetPathParam("id", problem).
		SetFormDataAnyType(map[string]interface{}{
			"runtime": runtime,
			"stream":  streamed,
		}).
		EnableForceMultipart().
		SetContext(ctx).
		SetFile("code", src)
	if streamed {
		req.DisableAutoReadResponse()
	} else {
		req.SetSuccessResult(&submissionId)
	}
	res, e := req.Post("/problems/{id}/submit")
	if e != nil {
		return nil, e
	}
	if !res.IsSuccessState() {
		var err common.Error
		buf, _e := io.ReadAll(res.Body)
		if _e != nil {
			return nil, types.ErrReadBody
		}
		if json.Unmarshal(buf, &err) != nil {
			return nil, errors.New(string(buf))
		}
		return nil, errors.New(err.Message)
	}
	if !streamed {
		return submissionId, nil
	} else {
		resChan := make(chan interface{}, 1)
		reader := json.NewDecoder(res.Body)
		var _res submission.ResultResponse
		go func() {
			defer close(resChan)
			for reader.More() {
				if reader.Decode(&_res) != nil {
					return
				}
				switch _res.Type {
				case submission.ResponseTypeMetadata:
					var meta submission.Metadata
					if err := mapstructure.Decode(_res.Data, &meta); err != nil {
						c.l.Fatal("error decoding metadata: %s", err)
					}
					resChan <- meta
				case submission.ResponseTypeAck:
					resChan <- nil
				case submission.ResponseTypeCase:
					var cr submission.CaseResult
					if err := mapstructure.Decode(_res.Data, &cr); err != nil {
						c.l.Fatal("error decoding case response: %s", err)
					}
					resChan <- cr
				case submission.ResponseTypeFinal:
					var fr submission.FinalResult
					if err := mapstructure.Decode(_res.Data, &fr); err != nil {
						c.l.Fatal("error decoding final response: %s", err)
					}
					resChan <- fr
					return
				}
			}

		}()
		return resChan, nil
	}
}

func (c *Client) SubmitStreamed(ctx context.Context, problem, src, runtime string) (chan interface{}, error) {
	v, e := c.submitGeneric(ctx, problem, src, runtime, true)
	if e != nil {
		return nil, e
	}
	return v.(chan interface{}), nil
}

func (c *Client) Submit(ctx context.Context, problem, src, runtime string) (uint32, error) {
	v, e := c.submitGeneric(ctx, problem, src, runtime, false)
	if e != nil {
		return 0, e
	}
	return v.(uint32), nil
}
