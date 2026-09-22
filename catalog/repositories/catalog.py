from models.models import Brands
from sqlalchemy.ext.asyncio import AsyncSession
from fastapi import HTTPException, status

async def addCategories(data, db: AsyncSession):
    brand = Brands(
        name = data.name
    )
    db.add(brand)

    await db.commit()
    await db.refresh(brand)

    return {
        "id": brand.id
    }
