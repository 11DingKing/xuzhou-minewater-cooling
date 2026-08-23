package health
import "testing"
import "context"
func TestMinewater021(t *testing.T){if Check(context.Background(),nil).Healthy{t.Fatal("nil database reported healthy")}}