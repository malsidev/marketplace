from fastapi import UploadFile
from uuid import uuid4
from database.miniodb import client

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