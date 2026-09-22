from sqlalchemy import ForeignKey, String, TIMESTAMP, func
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship


class Base(DeclarativeBase):
    pass


class Products(Base):
    __tablename__ = "products"

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str] = mapped_column(String(55),nullable=False)
    description: Mapped[str] = mapped_column(String(255))
    brand_id: Mapped[int] = mapped_column(ForeignKey("brands.id"),nullable=False)
    category_id: Mapped[int] = mapped_column(ForeignKey("categories.id"),nullable=False)
    created_at: Mapped[int] = mapped_column(TIMESTAMP(timezone=True),server_default=func.now())

    brand: Mapped["Brands"] = relationship(back_populates="products")
    category: Mapped["Categories"] = relationship(back_populates="products")


class Categories(Base):
    __tablename__ = "categories"

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str] = mapped_column(String(55),nullable=False)

    products: Mapped[list["Products"]] = relationship(back_populates="category")


class Brands(Base):
    __tablename__ = "brands"

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str] = mapped_column(String(55),nullable=False)

    products: Mapped[list["Products"]] = relationship(back_populates="brand")