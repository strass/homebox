<div align="center">
   <svg width="200" align="center" viewBox="0 0 10817 9730" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:5.42683">
      <path d="M9310.16 2560.9c245.302 249.894 419.711 539.916 565.373 845.231 47.039 98.872 36.229 215.514-28.2 304.05-64.391 88.536-172.099 134.676-280.631 120.28 0 .053-.039.053-.039.053" style="fill:gray;stroke:#000;stroke-width:206.41px"/>
      <path d="M5401.56 487.044c-127.958 6.227-254.855 40.77-370.992 103.628-765.271 414.225-2397.45 1297.68-3193.03 1728.32-137.966 74.669-250.327 183.605-328.791 313.046l3963.09 2122.43s-249.048 416.428-470.593 786.926c-189.24 316.445-592.833 429.831-919.198 258.219l-2699.36-1419.32v2215.59c0 226.273 128.751 435.33 337.755 548.466 764.649 413.885 2620.97 1418.66 3385.59 1832.51 209.018 113.137 466.496 113.137 675.514 0 764.623-413.857 2620.94-1418.63 3385.59-1832.51 208.989-113.136 337.743-322.193 337.743-548.466v-3513.48c0-318.684-174.59-611.722-454.853-763.409-795.543-430.632-2427.75-1314.09-3193.02-1728.32-141.693-76.684-299.364-111.227-455.442-103.628" style="fill:#dadada;stroke:#000;stroke-width:206.42px"/>
      <path d="M5471.83 4754.46V504.71c-127.958 6.226-325.127 23.1-441.264 85.958-765.271 414.225-2397.45 1297.68-3193.03 1728.32-137.966 74.669-250.327 183.605-328.791 313.046l3963.09 2122.43Z" style="fill:gray;stroke:#000;stroke-width:206.42px"/>
      <path d="m1459.34 2725.96-373.791 715.667c-177.166 339.292-46.417 758 292.375 936.167l4.75 2.5m0 0 2699.37 1419.29c326.374 171.625 729.916 58.25 919.165-258.208 221.542-370.5 470.583-786.917 470.583-786.917l-3963.04-2122.42-2.167 3.458-47.25 90.458" style="fill:#dadada;stroke:#000;stroke-width:206.42px"/>
      <path d="M5443.74 520.879v4149.79" style="fill:none;stroke:#000;stroke-width:153.5px"/>
      <path d="M8951.41 4102.72c0-41.65-22.221-80.136-58.291-100.961-36.069-20.825-80.51-20.825-116.58 0l-2439.92 1408.69c-36.07 20.825-58.29 59.311-58.29 100.961V7058c0 41.65 22.22 80.136 58.29 100.961 36.07 20.825 80.51 20.825 116.58 0l2439.92-1408.69c36.07-20.825 58.291-59.312 58.291-100.962v-1546.59Z" style="fill:#567f67"/>
      <path d="M8951.41 4102.72c0-41.65-22.221-80.136-58.291-100.961-36.069-20.825-80.51-20.825-116.58 0l-2439.92 1408.69c-36.07 20.825-58.29 59.311-58.29 100.961V7058c0 41.65 22.22 80.136 58.29 100.961 36.07 20.825 80.51 20.825 116.58 0l2439.92-1408.69c36.07-20.825 58.291-59.312 58.291-100.962v-1546.59ZM6463.98 5551.29v1387.06l2301.77-1328.92V4222.37L6463.98 5551.29Z"/>
      <path d="M5443.76 9041.74v-4278.4" style="fill:none;stroke:#000;stroke-width:206.44px;stroke-linejoin:miter"/>
      <path d="m5471.79 4773.86 3829.35-2188.22" style="fill:none;stroke:#000;stroke-width:206.43px;stroke-linejoin:miter"/>
   </svg>
</div>
<h1 align="center" style="margin-top: -10px"> HomeBox </h1>
<p align="center" style="width: 100;">
   <a href="https://hay-kot.github.io/homebox/">Docs</a>
   |
   <a href="https://homebox.fly.dev">Demo</a>
   |
   <a href="https://discord.gg/tuncmNrE4z">Discord</a>
</p>

## Quick Start

```yml
version: "3.4"
 services:
   homebox:
     image: ghcr.io/hay-kot/homebox:nightly
     container_name: homebox
     restart: always
     volumes:
       - homebox-data:/data/
     ports:
       - 3100:7745

volumes:
   homebox-data:
     driver: local
```

## MVP Todo

- [ ] Asset Attachments for Items
- [ ] Db Migrations
  - [ ] How To
- [x] Documentation
  - [x] Docker Compose
  - [x] Config Options
- [x] Locations
  - [x] Create
  - [x] Update
  - [x] Delete
- [x] Labels
  - [x] Create
  - [x] Update
  - [x] Delete
- [x] Items CRUD
  - [x] Create
  - [x] Update
  - [x] Delete
- [x] Fields To Add
  - [x] Quantity
  - [x] Insured (bool)
- [x] Bulk Import via CSV
  - [x] Initial
  - [x] Add Warranty Columns
  - [x] All Fields
  - [x] Documentations
- [x] Release Flow
  - [x] CI/CD Docker Builds w/ Multi-arch
  - [x] Auto Fly.io Deploy for Nightly
  - [x] Deploy Docs
- [x] Repo House Keeping
  - [x] Add License
  - [x] Issues Template
  - [x] PR Templates
  - [x] Security Policy
  - [x] Feature Request Template
- [x] Embedded Version Info
  - [x] Version Number
  - [x] Git Hash
- [x] Setup Docker Volumes in Dockerfile
- [x] Warranty Information
  - [x] Option for Lifetime Warranty or Warranty Period

## All Todo's

- [ ] Dev Container for Development
- [ ] User Invitation Links to Join Group
- [ ] Maintenance Logs
  - [ ] Schedule Future Maintenance
  - [ ] Email on Maintenance Due
  - [ ] QR Code Stickers to Scan to enter a Maintenance Task
- [ ] Export CSV (With IDs)
- [ ] User Profile
  - [ ] Adjust Theme (Daisy UI)
  - [ ] Delete Profile
  - [ ] Send User Invites
  - [ ] Set Currency
  - [ ] Change Password
- [ ] Admin Page
  - [ ] Instance Statistics
  - [ ] User Management
    - [ ] Delete User
    - [ ] Reset Password
- [ ] Expose Swagger API Documentation
  - [ ] Dynamic Port / Host Settings

## Credits

- Logo by [@lakotelman](https://github.com/lakotelman)
