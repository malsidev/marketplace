from fastapi import APIRouter

router = APIRouter(prefix="/catalog", tags=["catalog"])

@router.get("/")
async def get_catalog():
    return {"message": "Welcome to the catalog!"}