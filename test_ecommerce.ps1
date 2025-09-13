# -------------------------
# E-commerce Backend Test
# -------------------------

# 1️⃣ Login ve token alma
$loginHeaders = @{ "Content-Type" = "application/json" }
$loginBody = '{ "username": "admin", "password": "123456" }'
$response = Invoke-RestMethod -Uri http://localhost:8080/login -Method POST -Headers $loginHeaders -Body $loginBody

$token = $response.access_token
$authHeaders = @{ "Authorization" = "Bearer $token"; "Content-Type" = "application/json" }

Write-Host "✅ Login successful. Access token acquired." -ForegroundColor Green

# 2️⃣ Tüm ürünleri listele
Write-Host "`n📋 Tüm ürünler:"
$allProducts = Invoke-RestMethod -Uri http://localhost:8080/products -Method GET -Headers $authHeaders
$allProducts | Format-Table id, name, description, price, quantity

# 3️⃣ Yeni ürün ekle
Write-Host "`n➕ Yeni ürün ekleme..."
$newProductBody = '{ "name": "Mouse", "description": "Wireless Mouse", "price": 25.5, "quantity": 50 }'
Invoke-RestMethod -Uri http://localhost:8080/products -Method POST -Headers $authHeaders -Body $newProductBody
Write-Host "✅ Yeni ürün eklendi."

# 4️⃣ Ürünü güncelle (örnek: ID 1)
Write-Host "`n✏️ Ürün güncelleme (ID 1)..."
$updateProductBody = '{ "name": "Laptop Pro", "description": "Gaming Laptop Updated", "price": 1300.0, "quantity": 8 }'
Invoke-RestMethod -Uri http://localhost:8080/products/1 -Method PUT -Headers $authHeaders -Body $updateProductBody
Write-Host "✅ Ürün güncellendi."

# 5️⃣ Ürünü sil (örnek: ID 3)
Write-Host "`n🗑️ Ürün silme (ID 3)..."
Invoke-RestMethod -Uri http://localhost:8080/products/3 -Method DELETE -Headers $authHeaders
Write-Host "✅ Ürün silindi."

# 6️⃣ Tek ürün çek (örnek: ID 2)
Write-Host "`n🔎 Tek ürün çekme (ID 2)..."
$product = Invoke-RestMethod -Uri http://localhost:8080/products/2 -Method GET -Headers $authHeaders
$product | Format-Table id, name, description, price, quantity
