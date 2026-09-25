from minio import Minio
from minio.error import S3Error

client = Minio(
    "localhost:9000",
    access_key="minioadmin",
    secret_key="minioadmin",
    secure=False
)
BUCKET_NAME = "products"

def create_bucket_if_not_exists():
    try:
        if not client.bucket_exists(BUCKET_NAME):
            client.make_bucket(BUCKET_NAME)
            print(f"Bucket '{BUCKET_NAME}' created")
        else:
            print(f"Bucket '{BUCKET_NAME}' already exists")

    except S3Error as e:
        print(f"MinIO error: {e}")
        raise