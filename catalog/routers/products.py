from fastapi import APIRouter, Header, UploadFile, File

from services.products import postProducts, postProductsImages
from schemas.schemas import Products

router = APIRouter(prefix="/products", tags=["products"])

@router.post("")
async def post_products(
        product_data : Products,
        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await postProducts(user_id)

@router.post("/{product_id}/images")
async def post_products_images(product_id : int, image: UploadFile = File(...)):
    return await postProductsImages(product_id, image)