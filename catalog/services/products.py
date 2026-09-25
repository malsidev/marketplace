from repositories.product import addProduct
from services.minio import upload_image

async def postProducts(name, description, brand_id, category_id, image, db):
    product_id = await addProduct(name, description, brand_id, category_id, db)
    await upload_image(image, product_id)
    return {"message" : "ok"}

async def postProductsImages(product_id, image):
    return await upload_image(image, product_id)