package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gorm.io/gorm"
)

const (
	Sender   = "info@mespitech.com"
	Subject  = "mespitech - Staging Subject"
	HtmlBody = "<div><h2>%v</h2><p>IDコードは次の通りです: %v</p><br/><h2>コードを共有しないでください。宜しくお願い致します。!</h2></div>"

	TextBody = "This is test email(staging)."
	CharSet  = "UTF-8"
)

type EmailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *EmailRepository {
	return &EmailRepository{
		db: db,
	}
}

func (r EmailRepository) SendOTPForRegister(email string, otpCode string) error {
	// Create new connection
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return err
	}

	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(fmt.Sprintf("OTP : %v", otpCode)),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	// Send email
	_, err = svc.SendEmail(input)
	if err != nil {
		return err
	}
	return nil
}
