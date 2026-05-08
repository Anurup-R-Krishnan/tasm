export interface BaseEntity {
  id: number;
  createdAt: string;
  updatedAt: string;
}

export interface Asset extends BaseEntity {
  name: string;
  tagId: string;
  category: string;
  location: string;
  status: string;
  custodian: string;
  currentStock: number;
  reorderLevel: number;
  value: number;
  purchaseDate: string;
  warrantyStatus: string;
  warrantyExpiry: string;
}

export interface AuditSession extends BaseEntity {
  title: string;
  startDate: string;
  endDate?: string;
  status: string;
  totalAssets: number;
  scannedAssets: number;
  discrepancyCount: number;
  auditorName: string;
  progress: number;
}

export interface AuditDiscrepancy extends BaseEntity {
  auditSessionId: number;
  assetName: string;
  assetTag: string;
  issueType: string;
  lastKnownLocation: string;
  scannedLocation: string;
  recommendedAction: string;
  status: string; // Open, Resolved, Dismissed
  resolution: string;
  resolvedBy: string;
}

export interface Consumable extends BaseEntity {
  name: string;
  category: string;
  currentStock: number;
  reorderLevel: number;
  location: string;
}

export interface MaintenanceContract extends BaseEntity {
  vendorName: string;
  serviceType: string;
  startDate: string;
  endDate: string;
  status: string;
}

export interface WorkOrder extends BaseEntity {
  workOrderId: string;
  title: string;
  assetLocation: string;
  assetName?: string;
  assetTag?: string;
  description?: string;
  issue: string;
  severity: string;
  reportedBy?: string;
  targetDate: string;
  status: string;
  technician: string;
  cost: number;
  estimatedCost?: number;
}

export interface ProcurementRequest extends BaseEntity {
  title: string;
  status: string;
  priority: string;
  estimatedValue: number;
  actualValue: number;
  requestorName: string;
  requestorInitials: string;
  department: string;
  poNumber: string;
}

export interface LedgerEntry extends BaseEntity {
  transactionId: string;
  date: string;
  description: string;
  amount: number;
  type: string;
  category: string;
}

export interface LeaseAgreement extends BaseEntity {
  leaseId: string;
  vendor: string;
  startDate: string;
  endDate: string;
  monthlyCost: number;
  status: string;
}

export interface DepreciationSchedule extends BaseEntity {
  assetId: string;
  assetName: string;
  purchaseValue: number;
  currentValue: number;
  method: string;
}

export interface SoftwareLicense extends BaseEntity {
  softwareName: string;
  planName: string;
  status: string;
  totalSeats: number;
  usedSeats: number;
  renewalDate: string;
  annualCost: number;
}

export interface SystemUser extends BaseEntity {
  employeeId: string;
  name: string;
  email: string;
  department: string;
  role: string;
  status: string;
  lastLogin: string;
}

export interface Reservation extends BaseEntity {
  title: string;
  type: string;
  status: string;
  startTime: string;
  endTime: string;
  location: string;
  bookedBy: string;
}

export interface Location extends BaseEntity {
  name: string;
  type: string;
  address: string;
  capacity: number;
  status: string;
}

export interface UserRole extends BaseEntity {
  roleName: string;
  description: string;
  usersCount: number;
}

export interface SystemAlert extends BaseEntity {
  title: string;
  message: string;
  type: string;
  source: string;
  isRead: boolean;
}

export interface AssetEvent extends BaseEntity {
  assetId: number;
  eventType: string;
  actorID?: number;
  actorName?: string;
  description?: string;
  previousStatus?: string;
  newStatus?: string;
  previousCustodian?: string;
  newCustodian?: string;
  dueDate?: string;
  notes?: string;
}

export interface ScanResult {
  result: 'found' | 'unregistered' | 'error';
  asset?: Asset;
  locationMatch: boolean;
  progress?: number;
  discrepancyId?: number;
}
