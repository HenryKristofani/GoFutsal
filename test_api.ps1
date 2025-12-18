# Script untuk test API GoFutsal
Write-Host "Testing GoFutsal API..."

# Test health endpoint
Write-Host "1. Testing health endpoint..."
try {
    $health = Invoke-WebRequest -Uri "http://localhost:8080/health" -Method GET
    Write-Host "Health check: $($health.StatusCode) - $($health.Content)"
} catch {
    Write-Host "Health check failed: $($_.Exception.Message)"
}

# Register admin user
Write-Host "2. Registering admin user..."
$adminUser = @{
    username = "admin"
    email = "admin@gofutsal.com"
    password = "admin123"
    role = "admin"
} | ConvertTo-Json

try {
    $registerResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/users/register" -Method POST -Body $adminUser -ContentType "application/json"
    Write-Host "Register response: $($registerResponse.StatusCode) - $($registerResponse.Content)"
} catch {
    Write-Host "Register failed: $($_.Exception.Message)"
}

# Test login
Write-Host "3. Testing login..."
$loginData = @{
    username = "admin"
    password = "admin123"
} | ConvertTo-Json

try {
    $loginResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/auth/login" -Method POST -Body $loginData -ContentType "application/json"
    Write-Host "Login response: $($loginResponse.StatusCode) - $($loginResponse.Content)"
} catch {
    Write-Host "Login failed: $($_.Exception.Message)"
}

Write-Host "API test completed!"