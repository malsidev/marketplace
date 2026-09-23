from services.minio import upload_image

async def postProducts(data):
    pass

async def postProductsImages(product_id, image):
    return await upload_image(image, product_id)