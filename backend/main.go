import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"log"
)
func main(){
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	svc := dynamodb.New(sess)
	tableName:="users"
	input:=&dynamodb.CreateTableInput{
		AttributeDefinitions:[]*dynamodb.AttributeDefinitions{
			AttributeName:aws.String("username"),
			AttributeType:aws.String("S")
		},
		{
			AttributeName:aws.String("password"),
			AttributeType:aws.String("S")
		}
	}
}