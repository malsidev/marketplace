from fastapi import APIRouter, Header, UploadFile, File, Form, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from database.pgdb import get_db
from services.products import postProducts, postProductsImages
from schemas.schemas import Products

router = APIRouter(prefix="/products", tags=["products"])

@router.post("")
async def post_products(
        name: str = Form(...),
        description: str = Form(...),
        brand_id: int = Form(...),
        category_id: int = Form(...),
        image: UploadFile = File(...),
        db: AsyncSession = Depends(get_db),

        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await postProducts(name, description, brand_id, category_id, image, db)

@router.post("/{product_id}/images")
async def post_products_images(product_id : int, images: UploadFile = File(...)):
    return await postProductsImages(product_id,  images)