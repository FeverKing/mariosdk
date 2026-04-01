package sdkclient

import (
	"testing"

	"github.com/FeverKing/mariosdk/sdk/sdkreq"
)

// 校验 SDK 能反序列化比赛题库中的理论题导出字段。
func TestGetProblemBankForCompetitionDecodesTheoryFields(t *testing.T) {
	client := newAuthedClient(&fakeRequester{
		respBody: `{"code":200,"msg":"ok","data":{"id":"9","name":"bank-9","des":"d9","problemNum":3,"problems":[{"id":"2001","name":"choice-1","desc":"choice-desc","problemType":3,"difficulty":1,"options":"A,B,C,D","answers":"A,C","score":100},{"id":"2002","name":"fill-1","desc":"fill-desc","problemType":4,"difficulty":2,"answers":"flag{fill}","score":100},{"id":"2003","name":"subjective-1","desc":"subjective-desc","problemType":7,"difficulty":3,"referenceAnswer":"ref answer","reviewGuideline":"score by key points","score":120}]}}`,
	})

	resp, err := client.GetProblemBankForCompetition(&sdkreq.GetProblemBankForCompetitionReq{})
	if err != nil {
		t.Fatalf("GetProblemBankForCompetition() error = %v", err)
	}
	if len(resp.Problems) != 3 {
		t.Fatalf("expected 3 problems, got %d", len(resp.Problems))
	}
	if resp.Problems[0].Options != "A,B,C,D" || resp.Problems[0].Answers != "A,C" || resp.Problems[0].Score != 100 {
		t.Fatalf("choice theory fields not decoded: %+v", resp.Problems[0])
	}
	if resp.Problems[1].Answers != "flag{fill}" || resp.Problems[1].Score != 100 {
		t.Fatalf("fill-blank theory fields not decoded: %+v", resp.Problems[1])
	}
	if resp.Problems[2].ReferenceAnswer != "ref answer" || resp.Problems[2].ReviewGuideline != "score by key points" || resp.Problems[2].Score != 120 {
		t.Fatalf("subjective theory fields not decoded: %+v", resp.Problems[2])
	}
}
