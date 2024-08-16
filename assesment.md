# Examen Teórico - Riwi Coder Ruta NodeJS - NestJS

**Duración:** 2 horas y 15 minutos  
**Número de preguntas:** 40  
**Hora de inicio:** 7:10 PM  
**Hora de finalización:** 9:25 PM  
**Instrucciones de entrega:**  
- **Formato de entrega:** Archivo `.txt` con las respuestas y un archivo `PNG` de los diagramas de Draw.io comprimidos en un archivo `nombre-apellido-ultimos-4-digitos-cedula.zip`
- **Plazo de entrega:** Hasta las 9:30 PM, carga a Moodle.

---

0. **Sondeo de horario: Tienes disponibilidad para un cambio de jornada sin cambio de ruta (misma ruta NodeJS - NestJS)?**
   a) Sí  
   b) No

### **Sección 1: Lenguajes - JavaScript**

**Introducción y Standards (3 preguntas)**

1. **¿Cuál de las siguientes es una buena práctica para nombrar variables en JavaScript?**  
   a) Utilizar caracteres especiales  
   b) Comenzar con un número  
   c) Usar camelCase  
   d) Usar palabras reservadas del lenguaje  
   **(Selección única)**

2. **¿Qué paradigma de programación es principalmente soportado por JavaScript?**  
   a) Funcional  
   b) Orientado a objetos  
   c) Imperativo  
   d) Declarativo  
   **(Selección única)**

3. **¿Cuál de las siguientes es una declaración válida de clase en JavaScript?**  
   a) `class MyClass extends ParentClass {}`  
   b) `function MyClass() {}`  
   c) `const MyClass = new class {}`  
   d) `var MyClass = function() {}`  
   **(Selección única)**

**Variables (2 preguntas)**

4. **¿Cuál es la principal diferencia entre `let` y `const` en JavaScript?**  
   a) `let` permite reasignación, `const` no  
   b) `let` es solo para números, `const` para strings  
   c) `let` es global, `const` es local  
   d) `let` es solo para funciones, `const` para objetos  
   **(Selección única)**

5. **¿Cuál de las siguientes opciones no es una forma válida de declarar variables en JavaScript?**  
   a) `var`  
   b) `let`  
   c) `define`  
   d) `const`  
   **(Selección única)**

**Tipos de Datos (2 preguntas)**

6. **¿Cuál de los siguientes es un tipo de dato primitivo en JavaScript?**  
   a) Object  
   b) Array  
   c) String  
   d) Function  
   **(Selección única)**

7. **¿Qué diferencia principal existe entre `null` y `undefined` en JavaScript?**  
   a) `null` indica ausencia de valor, `undefined` indica que una variable no ha sido asignada  
   b) `null` es un objeto, `undefined` es un tipo de dato  
   c) `null` se asigna por defecto, `undefined` es un valor constante  
   d) `null` y `undefined` son equivalentes  
   **(Selección única)**

**Truthy y Falsy (1 pregunta)**

8. **¿Cuál de los siguientes valores es considerado `falsy` en JavaScript?**  
   a) `0`  
   b) `1`  
   c) `"false"`  
   d) `[]`  
   **(Selección única)**

**Operadores (2 preguntas)**

9. **¿Cuál de los siguientes operadores se utiliza para comparar tanto el valor como el tipo en JavaScript?**  
   a) `==`  
   b) `===`  
   c) `!=`  
   d) `!==`  
   **(Selección única)**

10. **¿Qué resultado devolverá la expresión `2 + 3 * 4 / 2` en JavaScript?**  
    a) 10  
    b) 12  
    c) 14  
    d) 8  
    **(Selección única)**

**Strings y Métodos (2 preguntas)**

11. **¿Cuál de los siguientes métodos de string se utiliza para extraer una parte de un string en JavaScript?**  
    a) `slice()`  
    b) `splice()`  
    c) `split()`  
    d) `join()`  
    **(Selección única)**

12. **¿Cuál es la forma correcta de concatenar los strings `Hello` y `World` en JavaScript?**  
    a) `"Hello".concat("World")`  
    b) `Hello + World`  
    c) `Hello.concat(World)`  
    d) `Hello & World`  
    **(Selección única)**

**Arrays y Métodos (2 preguntas)**

13. **¿Cuál de los siguientes métodos agrega un elemento al final de un array en JavaScript?**  
    a) `push()`  
    b) `pop()`  
    c) `shift()`  
    d) `unshift()`  
    **(Selección única)**

14. **¿Cuál es la diferencia principal entre `map()` y `forEach()` en JavaScript?**  
    a) `map()` modifica el array original, `forEach()` no  
    b) `forEach()` retorna un nuevo array, `map()` no  
    c) `map()` retorna un nuevo array, `forEach()` no  
    d) `forEach()` puede ser encadenado, `map()` no  
    **(Selección única)**

**Objetos (2 preguntas)**

15. **¿Cuál es la forma correcta de acceder a la propiedad `name` de un objeto `person` en JavaScript?**  
    a) `person[name]`  
    b) `person.name`  
    c) `person->name`  
    d) `person::name`  
    **(Selección única)**

16. **¿Qué técnica se utiliza comúnmente para clonar un objeto en JavaScript?**  
    a) `Object.assign({}, obj)`  
    b) `obj.copy()`  
    c) `new Object(obj)`  
    d) `obj.clone()`  
    **(Selección única)**

**Estructuras de Control (2 preguntas)**

17. **¿Cuál es la sintaxis correcta para una estructura `if-else` en JavaScript?**  
    a) `if (condition) then { } else { }`  
    b) `if (condition) { } else { }`  
    c) `if condition then { } else { }`  
    d) `if condition { } else { }`  
    **(Selección única)**

18. **¿Qué palabra clave se utiliza para definir múltiples casos en una estructura `switch` en JavaScript?**  
    a) `case`  
    b) `option`  
    c) `when`  
    d) `elseif`  
    **(Selección única)**

**Funciones (2 preguntas)**

19. **¿Cuál de las siguientes es una forma válida de declarar una función en JavaScript?**  
    a) `function myFunc() { }`  
    b) `myFunc = function { }`  
    c) `let myFunc() = { }`  
    d) `declare myFunc() { }`  
    **(Selección única)**

20. **¿Cuál es la diferencia principal entre una arrow function y una función tradicional en JavaScript?**  
    a) La arrow function tiene su propio `this`  
    b) La arrow function no tiene su propio `this`  
    c) La arrow function solo puede ser usada dentro de clases  
    d) La arrow function requiere un nombre  
    **(Selección única)**

**Expresiones Regulares (1 pregunta)**

21. **¿Cuál de los siguientes símbolos se utiliza para denotar el inicio de una cadena en una expresión regular?**  
    a) `$`  
    b) `^`  
    c) `*`  
    d) `.`  
    **(Selección única)**

**Scoping y Hoisting (2 preguntas)**

22. **¿Qué es el hoisting en JavaScript?**  
    a) Un error común  
    b) El proceso de mover declaraciones al inicio del scope  
    c) La eliminación de variables no usadas  
    d) La creación de variables globales automáticamente  
    **(Selección única)**

23. **¿Cuál de los siguientes scopes es creado cuando se declara una variable dentro de una función en JavaScript?**  
    a) Scope de bloque  
    b) Scope local  
    c) Scope global  
    d) Scope condicional  
    **(Selección única)**

**Event Loop en JavaScript (2 preguntas)**

24. **¿En qué fase del Event Loop se ejecutan los callbacks de las funciones `setTimeout` en JavaScript?**  
    a) Poll  
    b) Timers  
    c) Idle  
    d) Check  
    **(Selección única)**

25. **¿Qué tarea tiene la prioridad más baja en el Event Loop de JavaScript?**  
    a) Callbacks de `setTimeout`  
    b) Callbacks de `setImmediate`  
    c) IO operations  
    d) Close callbacks  
    **(Selección única)**

**Promesas y Async/Await (2 preguntas)**

26. **¿Cuál es la sintaxis correcta para manejar errores en una función `async/await`?**  
    a) `try { await func() } catch (error) { }`  
    b) `await func() try { } catch (error) { }`  
    c) `try { await func() } finally { }`  
    d) `await func().catch(error => { })`  
    **(Selección única)**

27. **¿Qué retorna una función `async` en JavaScript?**  
    a) Una Promesa  
    b) Un Array  
    c) Un Objeto  
    d) Una Cadena  
    **(Selección única)**

**Prototipos (1 pregunta)**

28. **¿Qué es la herencia prototípica en JavaScript?**  
    a) Un modelo de herencia basado en objetos  
    b) La creación de una nueva clase desde una clase existente  
    c) La copia de un objeto en otro  
    d) La eliminación de propiedades de un objeto  
    **(Selección única)**

**Objeto `this` (1 pregunta)**

29. **¿Cuál es el valor de `this` en una función declarada llamada dentro de un objeto?**  
    a) El objeto que contiene la función  
    b) La función misma  
    c) El objeto global  
    d) `null`  
    **(Selección única)**

**POO: Clases y Métodos Estáticos (2 preguntas)**

30. **¿Cómo se declara un método estático en una clase en JavaScript?**  
    a) `static methodName() { }`  
    b) `methodName() static { }`  
    c) `function static methodName() { }`  
    d) `methodName static() { }`  
    **(Selección única)**

31. **¿Cuál de los siguientes es un pilar de la Programación Orientada a Objetos (POO)?**  
    a) Encapsulamiento  
    b) Flexibilidad  
    c) Recursión  
    d) Iteración  
    **(Selección única)**

**DOM (2 preguntas)**

32. **¿Cuál es la forma correcta de seleccionar un elemento con el id `myElement` en el DOM?**  
    a) `document.getElementById("myElement")`  
    b) `document.querySelector(".myElement")`  
    c) `document.querySelector("#myElement")`  
    d) `document.getElementById(.myElement)`  
    **(Selección única)**

33. **¿Cuál es la forma correcta de agregar un evento `click` a un botón en el DOM?**  
    a) `button.onClick = function() { }`  
    b) `button.addEventListener("click", function() { })`  
    c) `button.click = function() { }`  
    d) `button.addEvent("click", function() { })`  
    **(Selección única)**

**Módulos en JavaScript (1 pregunta)**

34. **¿Cuál es la forma correcta de exportar una función llamada `myFunction` en un módulo JavaScript?**  
    a) `export myFunction`  
    b) `export function myFunction() { }`  
    c) `export default myFunction()`  
    d) `function export myFunction() { }`  
    **(Selección única)**

**JSON-server (1 pregunta)**

35. **¿Qué comando se utiliza para iniciar un servidor con `json-server` en la terminal? (suponiendo que el archivo `db.json` existe y además tienes json-server instalado)**  
    a) `json-server --watch db.json`  
    b) `json-server start db.json`  
    c) `json-server init db.json`  
    d) `json-server run db.json`  
    **(Selección única)**

**JavaScript WebAPIs (1 pregunta)**

36. **¿Cuál es el método correcto para hacer una petición GET utilizando `fetch` en JavaScript?**  
    a) `fetch(url, { method: "POST" })`  
    b) `fetch(url, { method: "GET" })`  
    c) `fetch(url)`  
    d) `fetch(url).get()`  
    **(Selección única)**

**Paradigmas (1 pregunta)**

37. **¿Cuál de los siguientes paradigmas no es nativamente soportado por JavaScript?**  
    a) Imperativo  
    b) Funcional  
    c) Orientado a objetos  
    d) Lógico  
    **(Selección única)**

**Las siguientes preguntas evalúan habilidades de pensamiento crítico y resolución de problemas:**

**Pregunta 38: Análisis y Reflexión (Metacognición y Observación Reflexiva)**
**Enunciado:**  
Supón que estás desarrollando una aplicación web y has implementado una funcionalidad que, a pesar de estar técnicamente correcta, no proporciona la experiencia de usuario esperada. Reflexiona sobre el proceso que seguiste para implementar esta funcionalidad:

1. ¿Qué pasos seguiste para identificar la raíz del problema?  
2. ¿Qué errores podrías haber anticipado al comenzar esta tarea?  
3. ¿Cómo podrías haber abordado el problema de manera diferente desde el principio para evitar este resultado?  
4. Después de reflexionar, ¿qué cambios implementarías para mejorar esta funcionalidad?

**Pregunta 39: Resolución de Problemas Complejos (Experimentación Activa)**
**Enunciado:**  
Te enfrentas a un error recurrente en una aplicación web que impide la correcta integración con una API externa. Después de varios intentos fallidos de resolverlo, decides abordar el problema desde una nueva perspectiva. 

1. Propón una estrategia diferente para solucionar el error, incluyendo al menos tres enfoques alternativos.
2. ¿Cómo aplicarías cada uno de estos enfoques para determinar cuál es el más efectivo?
3. ¿Qué método utilizarías para evaluar la efectividad de la solución implementada?

**Pregunta 40: Innovación y Aplicación del Conocimiento (Conceptualización Abstracta)**
**Enunciado:**
Imagina que estás encargado de optimizar una función en JavaScript que realiza múltiples operaciones asíncronas. Aunque la función funciona correctamente, su rendimiento es subóptimo. Después de analizar el código, identificas que la implementación actual no es eficiente y podría mejorarse significativamente.

**Código Original:**

```javascript
async function processData() {
  const result1 = await fetchData1();
  const result2 = await fetchData2(result1);
  const result3 = await fetchData3(result2);
  return result3;
}
```

Este código realiza las operaciones de manera secuencial, lo cual puede ser ineficiente.

### Tareas:

#### Optimización de la Implementación:

Propón una nueva implementación que optimice el rendimiento del código, explicando cómo y por qué los cambios sugeridos mejorarán la eficiencia considerando las operaciones asíncronas. Puedes buscar inspiración en conceptos como Promesas, Async/Await, y otras técnicas de JavaScript para mejorar la eficiencia de la función.

#### Prueba Comparativa:
Realiza una prueba comparativa entre la implementación original y tu propuesta optimizada. Describe los resultados obtenidos, incluyendo tiempos de ejecución y cualquier otro criterio relevante, y explica cómo estos resultados determinan la efectividad de tu optimización.