# TASM Asset Management System Audit Report

## A) Executive Summary

1. **Current State**: TASM has a polished UI with minimal functional backend wiring; most modules display mock/static data rather than real database records.

2. **Asset Management Core**: CRUD operations for assets work correctly via API, but no check-in/out, custody history, or barcode scanning integration exists.

3. **Missing Critical Workflows**: Check-in/out, maintenance scheduling notifications, audit discrepancy resolution, and alerts are UI facades without business logic.

4. **QR Scanner**: Component exists in `src/components/QRScanner.vue` but isn't integrated into actual audit workflows.

5. **Financial Trackers**: Ledgers, depreciation, leases, and licenses have CRUD but no calculation, reporting, or alerting functionality.

6. **Reservation System**: Only basic CRUD with overlap detection; no availability checking or integration with asset status.

7. **User Roles**: UI shows role management but no RBAC enforcement in backend handlers.

8. **Recommendation**: 60-80 days of development needed to achieve AssetTiger parity; Tier 1 features require immediate implementation.

---

## B) Feature Parity Matrix

| AssetTiger Feature | Expected Behavior | Current App Status | Evidence |
|---|---|---|---|
| Asset Tracking | Track asset location, custodian, status, history | **Partially** - CRUD exists, no custody history/timeline | `/api/assets` works, but no `events` table |
| Barcode Scanning | Mobile scanning, instant updates, audit registration | **Missing** - QR component exists but not integrated | `QRScanner.vue` unused in audit flow |
| Check-in/Check-out | Record custody transfers, due dates, history | **Missing** - No checkout endpoints or workflows | `AssetCheckOutFlow.vue` static, no `/checkout` API |
| Maintenance Scheduling | Schedule, track, alert on maintenance due dates | **Partially** - CRUD for contracts/orders only | No alerting, no recurrence logic |
| Audit History | Audit sessions, discrepancy reconciliation | **Partially** - Audit sessions & discrepancies CRUD | No reconciliation workflow, no status updates |
| Custom Reports | Build/save/export custom reports | **Missing** - Static builder UI | No `/reports/generate` endpoint |
| Email Alerts | Configurable alerts for expirations, maintenance | **Missing** - No alerting system | No background jobs, no alert creation |
| Roles & Permissions | RBAC with user-level access control | **Facade** - UI shows roles, no enforcement | All handlers skip authorization checks |
| Reservations | Reserve assets, check availability, prevent conflicts | **Partially** - CRUD with overlap detection | No asset availability linkage |
| Depreciation | Calculate, track, report asset depreciation | **Partially** - CRUD entries only | No calculations, no value tracking |
| License Tracking | Track software licenses, renewal dates, seats | **Partially** - CRUD exists | No expiring alerts |
| Inventory/Consumables | Track stock levels, reorder points | **Partially** - CRUD exists | No low-stock alerts |
| Mobile App | Dedicated mobile interface for field work | **Missing** - Desktop-only SPA | No PWA or mobile-specific views |
| Document Management | Attach files to assets, contracts | **Missing** - No file upload support | No storage integration |
| Custom Fields | Add custom metadata to assets | **Missing** - Fixed schema only | No JSONB or EAV pattern |

---

## C) User Journey Audit

### Complete Asset Lifecycle Flow

**1. Create**: ✅ Working
- Navigate to `/inventory` → Click "Add Asset" → Fill form → Submit
- API: POST `/api/assets` → Returns 201 with asset ID

**2. Tag**: ❌ Missing
- No barcode generation endpoint
- No QR code printing capability
- QR scanner component exists but detached from asset context

**3. Assign**: ⚠️ Limited
- Can set custodian field on create/edit
- No assignment history tracked
- No notification to custodian

**4. Check-out**: ❌ Broken
- UI flow at `/check-out` exists but collects no data
- No checkout endpoint (`/checkout`) exists
- Asset status doesn't update

**5. Return**: ❌ Broken
- No check-in endpoint
- No late return detection
- No condition update on return

**6. Audit**: ⚠️ Partial
- Audit sessions can be created (`POST /api/audits`)
- Discrepancies can be logged (`POST /api/discrepancies`)
- No reconciliation workflow to update asset status
- QR scanner exists but scanning doesn't update audit progress

**7. Maintain**: ⚠️ Facade
- Maintenance contracts CRUD works
- Work orders CRUD works
- No scheduling engine
- No notifications for due maintenance

**8. Dispose**: ❌ Missing
- No disposal workflow
- No disposal reason tracking
- No value write-off processing

---

## D) Technical Findings

### Backend Missing Validations
1. **No Unique Constraints**: TagID can be duplicated (no `UNIQUE` constraint)
2. **No Foreign Keys**: Asset references (custodian, location) aren't enforced
3. **No Soft Delete**: Assets are hard-deleted, losing history
4. **No Audit Trail**: No `events` or `audit_logs` table for change tracking
5. **No Input Sanitization**: XSS vectors in name/tag fields possible

### Broken/Missing API Routes
1. `/api/assets/:id/checkout` - 404
2. `/api/assets/:id/checkin` - 404
3. `/api/assets/:id/history` - 404
4. `/api/reports` - 404
5. `/api/reports/:id/export` - 404
6. `/api/alerts/create` - No route for automatic alerts

### Data Model Mismatches
1. **Asset Model**: No foreign key relationships (`Custodian` stored as string, not User ID reference)
2. **Reservation Model**: No link to Asset ID
3. **WorkOrder Model**: References asset by name/tag string, not relationship
4. **AuditDiscrepancy**: No link to AuditSession ID
5. **No Events Model**: Asset history must be tracked in separate events table

### Authorization Issues
1. All handlers use `requireDB(c)` but skip role/permission checks
2. JWT middleware exists but no RBAC middleware
3. Admin routes (delete, update) accessible to any authenticated user
4. No row-level security (users can see all assets regardless of role)

---

## E) Prioritized Backlog

### Tier 1: Critical (Make Product Purposeful)
1. **Check-in/out API** - Create `/api/assets/:id/checkout` and `/checkin` endpoints
2. **Asset Events History** - Add events table, log all asset status changes
3. **Role Enforcement** - Add authorization middleware to all handlers
4. **Audit Reconciliation** - Link discrepancies to audits, update asset status on resolve
5. **Barcode Integration** - Connect QR scanner to asset lookup and update

### Tier 2: Important Parity Gaps
1. **Alerting System** - Background job to create alerts for expiring items
2. **Maintenance Notifications** - Email alerts for due maintenance
3. **Reservation Availability** - Check asset status before allowing reservation
4. **Custom Reports** - Report builder that executes queries and exports data
5. **Document Upload** - File storage for asset attachments

### Tier 3: Enhancements
1. **Mobile PWA** - Make app installable on mobile devices
2. **Offline Mode** - Cache data for offline audit scanning
3. **Advanced Depreciation** - Automatic calculation engine
4. **Custom Fields** - EAV pattern for asset metadata
5. **Dashboard Analytics** - Real-time charts and metrics

---

## F) Recommendations

### Concrete Steps
1. **Week 1-2**: Implement asset events tracking and checkout/in endpoints
2. **Week 3-4**: Add authorization middleware, fix RBAC
3. **Week 5-6**: Build audit reconciliation workflow
4. **Week 7-8**: Implement alerting and notification system
5. **Week 9-10**: Connect QR scanner to asset workflows
6. **Week 11-12**: Build report generator and export functionality

### Metrics to Validate Functionality
- **Asset Check-out Success Rate**: Target >95%
- **Audit Completion Time**: Target <2 hours per 100 assets
- **Maintenance Compliance**: Target >90% on-time completion
- **Data Entry Accuracy**: Target <1% error rate
- **API Response Time**: Target <200ms average

### UX Adjustments
1. Replace "checked out" status display with actual checkout history
2. Add prominent "Check Out" button on asset detail page
3. Show overdue items in red on dashboard
4. Add barcode scanning entry point on mobile
5. Display next maintenance date on asset cards

---

## G) Evidence Appendix

### API Calls Tested (via curl)

**Authentication**
```bash
curl -X POST http://localhost/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@tasm.com","password":"admin123"}'
# Response: {"token":"eyJ...", "user":{"id":1,"name":"System Admin","role":"System Admin"}}
```

**Asset CRUD**
```bash
# Create
curl -X POST http://localhost/api/assets -H "Content-Type: application/json" -H "Authorization: Bearer <token>" \
  -d '{"name":"Test Asset","tagId":"T001","category":"IT Equipment","location":"Nila"}'
# Response: {"id":1,"name":"Test Asset","tagId":"T001",...}

# List
curl -H "Authorization: Bearer <token>" http://localhost/api/assets
# Response: [{"id":1,"name":"Test Asset",...}]
```

**Audit**
```bash
curl -H "Authorization: Bearer <token>" http://localhost/api/discrepancies
# Response: []  (empty - no data wired)
```

### UI Screens
- `http://localhost/` - Dashboard (SPA entry)
- `http://localhost/inventory` - Asset Registry (functional)
- `http://localhost/add-asset` - Add Asset Form (functional)
- `http://localhost/audit-scan` - QR Scanner (component mounted, no workflow)
- `http://localhost/check-out` - Check-out Flow (static UI)

### Backend Code References
- Asset handlers: `tasm-backend/handlers/asset_handler.go`
- Audit handlers: `tasm-backend/handlers/audit_handler.go`
- Maintenance handlers: `tasm-backend/handlers/maintenance_handler.go`
- Auth middleware: `tasm-backend/middleware/auth_middleware.go`

### Database Schema (from models/)
- `assets` - Main asset table (no events/history)
- `audit_sessions` - Audit session tracking
- `audit_discrepancies` - Discrepancy records
- `maintenance_contracts` - Vendor contracts
- `work_orders` - Maintenance tickets

---

**Conclusion**: TASM currently provides ~30% of AssetTiger's verified functionality. The polished UI masks a lack of critical business workflows. With focused development on Tier 1 items, the product could achieve baseline usability for basic asset tracking within 2-3 months.