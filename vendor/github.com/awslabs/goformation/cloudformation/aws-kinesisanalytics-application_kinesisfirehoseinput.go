package cloudformation

// AWSKinesisAnalyticsApplication_KinesisFirehoseInput AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.KinesisFirehoseInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisfirehoseinput.html
type AWSKinesisAnalyticsApplication_KinesisFirehoseInput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisfirehoseinput.html#cfn-kinesisanalytics-application-kinesisfirehoseinput-resourcearn
	ResourceARN *StringIntrinsic `json:"ResourceARN,omitempty"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisfirehoseinput.html#cfn-kinesisanalytics-application-kinesisfirehoseinput-rolearn
	RoleARN *StringIntrinsic `json:"RoleARN,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_KinesisFirehoseInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.KinesisFirehoseInput"
}
