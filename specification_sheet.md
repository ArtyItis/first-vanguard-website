# Specification Sheet

## Permission levels

- Siedler                             
- Offizier
- Konsul
- Inkasso
- Gouverneur
- Website-Admin

| Value | Title         |
| ----- | ------------- |
| 0     | Siedler       |
| 200   | Offizier      |
| 400   | Konsul        |
| 450   | Inkasso       |
| 600   | Gouverneur    |
| 1000  | Website-Admin |
## URL Parameters
| Error          | Type               | Success        | Application |
| -------------- | ------------------ | -------------- |------------ |
| authentication | login, company     | passwordChange | success     |
| login          | username, password | application    | failed      |
| password       | old, new           |
| user           | delete             |