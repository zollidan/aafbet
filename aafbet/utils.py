s3_client = boto3.client(
    's3',
    endpoint_url='https://storage.yandexcloud.net',
    aws_access_key_id=os.environ["AWS_ID"],
    aws_secret_access_key=os.environ["AWS_SECRET_KEY"]
)