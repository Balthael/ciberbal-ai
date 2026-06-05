---
description: Inicializa un nuevo engagement autorizado o auditoría mediana/grande usando Ciberbal en modo Hybrid-First para Cascade en Windsurf
---

# /sdd-new

Este workflow define el comportamiento obligatorio de **Cascade** al iniciar un nuevo engagement autorizado, auditoría, laboratorio, CTF o trabajo con incertidumbre suficiente como para requerir planificación formal.

## Propósito

Usar las capacidades nativas de Windsurf de forma **Hybrid-First**:

- **Plan Mode** para planificar
- **Memories / MCP (Engram)** para recuperar contexto previo
- **Artifacts `.sdd/`** solo como contrato formal de alcance, ROE y evidencia
- **Code Mode** únicamente después de aprobación explícita del usuario y dentro del alcance autorizado

## Cuándo usar este workflow

Activa este workflow cuando ocurra cualquiera de estas condiciones:

- El usuario inicia un **nuevo engagement autorizado**
- La tarea afecta **múltiples servicios, targets o artefactos de evidencia**
- El trabajo tiene **riesgo de alcance/ROE** o incertidumbre
- El usuario pide explícitamente trabajar con **Ciberbal** o el flujo legacy **SDD**
- La ejecución requiere un contrato formal antes de tocar targets o recolectar evidencia

Si la tarea es pequeña, puntual o claramente fuera de un engagement, este workflow NO es el camino correcto.

Antes de cualquier recon, explotación, validación o reporte, confirma que el target esté explícitamente autorizado: HTB/THM/lab/CTF, ROE firmado/autorización escrita, o auditoría defensiva sobre activos propios del usuario. Si la autorización no está clara, DETENTE y pregunta.

---

## Reglas operativas obligatorias

### 1. Cambiar inmediatamente a Plan Mode

Al comenzar este workflow, **DEBES entrar en Plan Mode de inmediato**.

Acciones obligatorias:

1. Analizar el pedido del usuario
2. Formular un plan de alto nivel
3. Identificar alcance, ROE, riesgos, dependencias y artefactos probables

Acciones prohibidas en esta etapa:

- NO ejecutar acciones contra targets
- NO entrar en Code Mode
- NO modificar sistemas ni evidencia
- NO ejecutar explotación parcial "para adelantar trabajo"
- NO asumir aprobación implícita

**Este workflow es de planificación formal, no de ejecución ofensiva.**

---

### 2. Recuperar contexto antes de proponer nada

Antes de redactar cualquier artefacto Ciberbal, **DEBES recuperar contexto del engagement, alcance, ROE y restricciones del proyecto**.

Orden de preferencia:

1. Usar **Engram** mediante las herramientas MCP canónicas: `mem_search` para buscar decisiones previas y `mem_context` para recuperar el contexto reciente del proyecto
2. Si Engram no está disponible o no devuelve contexto suficiente, leer `AGENTS.md`
3. Si existe contexto adicional del proyecto relacionado con Ciberbal, alcance, ROE o evidencia, incorporarlo también

Debes buscar, como mínimo:

- Decisiones previas del engagement o auditoría
- Convenciones del repositorio
- Restricciones de alcance/ROE
- Reglas de calidad o revisión
- Patrones ya establecidos para evidencia o reportes similares

Si no encuentras contexto suficiente, debes decirlo explícitamente en el plan. **No inventes convenciones.**

---

### 3. Crear el contrato formal inicial en `.sdd/`

Debes crear el directorio `.sdd/` si no existe.

Luego debes generar exactamente estos dos archivos iniciales:

- `.sdd/proposal.md`
- `.sdd/spec.md`

En esta fase, esos dos archivos son **obligatorios**.

#### Contenido mínimo de `.sdd/proposal.md`

Debe capturar, como mínimo:

- Título del engagement o auditoría
- Target/problema a investigar
- Objetivo
- Alcance incluido
- Alcance excluido
- ROE y límites de seguridad
- Enfoque propuesto
- Riesgos principales
- Supuestos abiertos
- Preguntas o decisiones pendientes

#### Contenido mínimo de `.sdd/spec.md`

Debe capturar, como mínimo:

- Requisitos de evidencia
- Requisitos de seguridad/ROE si aplican
- Escenarios de validación
- Criterios de aceptación de findings
- Restricciones técnicas relevantes del target/lab
- Casos límite conocidos o supuestos importantes

Los artefactos deben ser:

- Claros
- Revisables
- Ejecutables como contrato de evidencia y ejecución autorizada
- Consistentes con el contexto recuperado del proyecto

---

### 4. Presentar resumen de planificación al usuario

Después de crear `.sdd/proposal.md` y `.sdd/spec.md`, debes presentar un resumen breve y claro en chat.

Ese resumen debe incluir:

- Objetivo del engagement
- Alcance propuesto
- ROE y límites de seguridad
- Riesgos o dudas principales
- Confirmación de que los archivos fueron creados:
  - `.sdd/proposal.md`
  - `.sdd/spec.md`

No muestres una pared de texto innecesaria. Resume lo esencial para revisión.

---

### 5. Approval Gate absoluto

Una vez generados los documentos, debes **detenerte ABSOLUTAMENTE**.

Debes preguntar **exactamente**:

**¿Apruebas este plan de ejecución autorizada?**

Luego:

- Debes **esperar confirmación explícita**
- NO puedes continuar a Code Mode sin aprobación
- NO puedes empezar ejecución, explotación o recolección de evidencia "mientras tanto"
- NO puedes interpretar silencio como aprobación
- NO puedes reemplazar esta pausa con un resumen informal

Respuestas válidas para continuar:

- "sí"
- "aprobado"
- "dale"
- "go ahead"
- cualquier confirmación explícita equivalente

Si el usuario pide cambios:

- Debes seguir en Plan Mode
- Debes ajustar `.sdd/proposal.md` y/o `.sdd/spec.md`
- Debes volver a presentar el plan
- Debes volver a preguntar: **¿Apruebas este plan de ejecución autorizada?**

---

## Secuencia estricta de ejecución

Sigue esta secuencia sin saltos:

1. Detectar que el trabajo amerita `/sdd-new`
2. Entrar en **Plan Mode**
3. Recuperar contexto con **Engram** o, en su defecto, leer `AGENTS.md`
4. Sintetizar restricciones, alcance, ROE y riesgos
5. Crear `.sdd/` si no existe
6. Generar `.sdd/proposal.md`
7. Generar `.sdd/spec.md`
8. Presentar un resumen breve al usuario
9. Preguntar exactamente: **¿Apruebas este plan de ejecución autorizada?**
10. **Detenerte y esperar respuesta**

---

## Prohibiciones explícitas

Mientras este workflow no haya sido aprobado por el usuario:

- NO ejecutar acciones contra targets
- NO editar sistemas o evidencia fuera del plan
- NO ejecutar tareas de explotación/aplicación
- NO cambiar a Code Mode
- NO crear commits
- NO correr una ejecución parcial
- NO continuar automáticamente al siguiente paso de Ciberbal

---

## Criterio de salida de este workflow

Este workflow se considera correctamente ejecutado solo si:

- Cascade usó **Plan Mode**
- Recuperó contexto con **Engram** o `AGENTS.md`
- Generó `.sdd/proposal.md`
- Generó `.sdd/spec.md`
- Presentó un resumen al usuario
- Preguntó exactamente: **¿Apruebas este plan de ejecución autorizada?**
- Se detuvo a esperar aprobación explícita

Si cualquiera de esos puntos no ocurre, el workflow está mal ejecutado.
