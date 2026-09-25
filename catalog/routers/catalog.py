from fastapi import APIRouter, Depends, Header
from sqlalchemy.ext.asyncio import AsyncSession

from database import get_db
from services.catalog import getCatalog
from schemas.schemas import Products

router = APIRouter(prefix="/catalog", tags=["catalog"])

@router.get("")
async def get_catalog(
        db: AsyncSession = Depends(get_db),
        user_id: str | None = Header(default=None, alias="X-User-ID"),
):
    return await getCatalog(db)
