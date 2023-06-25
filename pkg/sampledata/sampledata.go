package sampledata

import "fmt"

const fileMain = `from fastapi import FastAPI
from fastapi.responses import RedirectResponse
from fastapi.middleware.cors import CORSMiddleware

from routes.products import product_router

import uvicorn

app = FastAPI()

# register origins

origins = ["*"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Register routes

app.include_router(product_router, prefix="/product")


@app.get("/")
async def home():
    return RedirectResponse(url="/docs")


if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8080, reload=True)

`

const fileModelsProduct = `from pydantic import BaseModel


class Product(BaseModel):
    id: int
    name: str
    description: str
    price: float
    stock: int

    class Config:
        schema_extra = {
            "example": {
                "id": 1,
                "name": "Jordan Retro 1",
                "description": "Iconic basketball sneakers with a classic design",
                "price": 199.99,
                "stock": 5
            }
        }

`

const fileRoutesProduct = `from typing import List
from fastapi import APIRouter, Body, HTTPException, Path, status

from models.products import Product

product_router = APIRouter(
    tags=["Product"]
)

products = []


@product_router.get("/", response_description="List all products", response_model=List[Product])
async def retrieve_all_products() -> List[Product]:
    return products


@product_router.get("/{id}", response_description="Get a single product")
async def retrieve_product(id: int = Path(..., title="The ID of the product to be recovered.")) -> Product:
    for product in products:
        if product.id == id:
            return product
    raise HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail="Product with supplied ID does not exist"
    )


@product_router.post("/", response_description="Add new product")
async def create_product(product: Product = Body(...)) -> dict:
    products.append(product)
    return {
        "message": "Product created successfully"
    }


@product_router.delete("/{id}", response_description="Delete product")
async def delete_product(id: int) -> dict:
    for product in products:
        if product.id == id:
            products.remove(product)
            return { "message": "Product deleted successfully" }
    raise HTTPException(
        status_code=status. HTTP_404_NOT_FOUND,
        detail="Event with supplied ID does not exist"
    )

`

const fileRequirements = `fastapi
uvicorn
`

const fileGitignore = `.venv
.env
__pycache__/

`

var fileFastmanconf = `# Fastman configuration file

dependencies_file: requirements.txt

scripts:
  install_package: python -m pip install {{package_name}}
  install_from_file: python -m pip install -r {{package_file}}
  uninstall_package: python -m pip uninstall {{package_name}}
  uninstall_from_file: python -m pip uninstall -r {{package_file}}
`

var sampleContent = map[string]string{
	"main.py":            fileMain,
	"models/products.py": fileModelsProduct,
	"routes/products.py": fileRoutesProduct,
	"requirements.txt":   fileRequirements,
	"fastmanconf.yaml":   fileFastmanconf,
}

func GetSampleContent(fileName string) (string, error) {
	content, ok := sampleContent[fileName]
	if !ok {
		return "", fmt.Errorf("file content does not exist: %s", fileName)
	}
	return content, nil
}
