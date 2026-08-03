# TODO - Fix InitPayment en core/services/flow.go

## Steps
- [x] Analizar el archivo core/services/flow.go y dependencias relacionadas
- [x] Identificar los errores de compilación con `go build ./...`
- [x] Cambiar `s.salesRepo.GetById(...)` → `s.salesRepo.Get(...)`
- [x] Eliminar variable `sale` no utilizada
- [x] Cambiar `s.cursoRepo.GetById(...)` → `s.cursoRepo.Get(...)` usando `cursoFilter`
- [x] Corregir llamada a `NewFlowService` en app/api/api.go (argumentos faltantes)
- [x] Verificar compilación con `go build ./...` ✅

