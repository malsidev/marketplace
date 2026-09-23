from models.models import Brands
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

async def addBrand(data, db: AsyncSession):
    brand = Brands(
        name = data.name
    )
    db.add(brand)

    await db.commit()
    await db.refresh(brand)

    return {
        "id": brand.id
    }

async def brand(db: AsyncSession):
    stmt = (select(Brands))

    res = await db.execute(stmt)
    data = res.scalars().all()

    return [
        {
        "id": brand.id,
        "name": brand.name
        }
        for brand in data
    ]