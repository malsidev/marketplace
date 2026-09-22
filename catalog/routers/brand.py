from fastapi import APIRouter, Header, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from database import get_db
from services.brand import brand
from schemas.schemas import Brands

router = APIRouter(prefix="/brand", tags=["prand"])

@router.post("")
async def post_brand(
        brand_data : Brands,
        db: AsyncSession = Depends(get_db),
        user_id: str | None = Header(default=None, alias="X-User-ID")
        
):
    return await brand(brand_data, db)