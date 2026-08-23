package approval
import "testing"
func TestMinewater005(t *testing.T){l:=New();_ = l.Create(Item{ID:"a",RegionID:"r",Applicant:"f",AmountCents:1});if e:=l.Transition("a",Paid,"u","");e==nil{t.Fatal("illegal transition")}}