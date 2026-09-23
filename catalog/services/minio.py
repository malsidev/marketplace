from minio import Minio
from minio.error import S3Error
from fastapi import UploadFile
from uuid import uuid4


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

async def upload_image(
    image: UploadFile,
    product_id: int,
):
    object_name = f"products/{product_id}/{uuid4()}.jpg"

    data = await image.read()

    from io import BytesIO

    client.put_object(
        "products",
        object_name,
        BytesIO(data),
        length=len(data),
        content_type=image.content_type,
    )

    return object_name