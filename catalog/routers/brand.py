from fastapi import APIRouter, Header, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from database.pgdb import get_db
from services.brand import postBrand, getBrand
from schemas.schemas import Brands

router = APIRouter(prefix="/brand", tags=["prand"])

@router.get("")
async def get_brand(db: AsyncSession = Depends(get_db), user_id: str | None = Header(default=None, alias="X-User-ID")):
    return await getBrand(db)

@router.post("")
async def post_brand(
        brand_data : Brands,
        db: AsyncSession = Depends(get_db),
        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await postBrand(brand_data, db)