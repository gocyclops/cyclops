package mys3

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

// InitS3 initializes the S3 client using credentials and endpoint information
// from environment variables. It retrieves the S3 URL, access key, and secret
// key from the environment, and uses them to configure the S3 client.
//
// Environment Variables:
// - S3_URL: The URL of the S3-compatible storage service (e.g., DigitalOcean Spaces).
// - S3_KEY_ID: The access key ID for the S3-compatible storage service.
// - S3_SECRET_ACCESS_KEY: The secret access key for the S3-compatible storage service.
//
// If any of the required environment variables are missing, the function logs an
// error message and returns without initializing the S3 client.
//
// The function uses static credentials and sets the region to "us-east-1". It
// initializes the S3 client with the provided credentials and endpoint, and sets
// the global S3Client variable to the initialized client.
func InitS3() {
	// Get the S3 URL from the environment variables
	spacesURL := os.Getenv("S3_URL")
	if spacesURL == "" {
		log.Println("Error, missing DigitalOcean Spaces URL in environment variables")
		return
	}

	// Get the AWS credentials from environment variables
	accessKey := os.Getenv("S3_KEY_ID")
	secretKey := os.Getenv("S3_SECRET_ACCESS_KEY")
	if accessKey == "" || secretKey == "" {
		log.Println("Error, missing Spaces credentials in environment variables")
		return
	}

	// Use static credentials provider
	staticCredentials := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))

	// Load the configuration with static credentials and region
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(staticCredentials),
	)
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize the S3 client with credentials and the DigitalOcean Spaces endpoint
	s3Client := s3.New(s3.Options{
		Credentials:      cfg.Credentials,
		Region:           cfg.Region,
		EndpointResolver: s3.EndpointResolverFromURL("https://" + spacesURL),
	})

	// Set the S3 client globally
	S3Client = s3Client
}

// DeleteFile deletes a file from an S3 bucket.
//
// Parameters:
//   - fileName: The name of the file to be deleted.
//
// Returns:
//   - error: An error object if the deletion fails, otherwise nil.
//
// The function constructs a DeleteObjectInput with the bucket name retrieved
// from the environment variable "S3_URL" and the provided file name. It then
// calls the DeleteObject method on the S3 client to delete the file. If the
// deletion fails, it logs the error and returns it.
func DeleteFile(fileName string) error {
	// Create the input for the delete request
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_URL")),
		Key:    aws.String(fileName),
	}

	_, err := S3Client.DeleteObject(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to delete file %s: %v", fileName, err)
		return err
	}

	return nil
}
