package query
import "testing"
func TestMinewater026(t *testing.T){q,a:=BuildWhere(Filter{RegionID:"r"});if q=="1=1"||len(a)==0{t.Fatal("missing region")}}