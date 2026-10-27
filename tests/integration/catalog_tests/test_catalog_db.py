from decimal import Decimal
from sqlalchemy import select
import pytest

from catalog.database.pgdb import AsyncSessionLocal
from catalog.models.models import Products, Brands, Categories


@pytest.mark.asyncio
async def test_create_and_read_product():
    async with AsyncSessionLocal() as session:
        brand = Brands(name="Test Brand")
        category = Categories(name="Test Category")

        session.add_all([brand, category])
        await session.flush()

        product = Products(
            name="Test Product",
            description="Integration test",
            brand_id=brand.id,
            category_id=category.id,
            price=Decimal("1500.00"),
        )

        session.add(product)
        await session.commit()

        product_id = product.id

    async with AsyncSessionLocal() as session:
        result = await session.execute(
            select(Products).where(Products.id == product_id)
        )
        saved_product = result.scalar_one()

        assert saved_product.name == "Test Product"
        assert saved_product.price == Decimal("1500.00")
        assert saved_product.brand_id is not None
        assert saved_product.category_id is not None