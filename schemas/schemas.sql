CREATE TABLE [KvBranch] (
  [AppId] int NOT NULL,
  [TenantId] int NOT NULL,
  [BranchId] int NOT NULL,
  [Name] nvarchar(4000) NULL,
  [BaseUtcOffsetMinute] int NULL,
  [IsActive] bit NOT NULL,
  [ModifiedDate] datetime2 NULL,
  [CreatedDate] datetime2 NOT NULL,
  [TimeZoneId] nvarchar(500) NULL
);

CREATE TABLE [Shift] (
  [Id] bigint NOT NULL,
  [Name] nvarchar(250) NOT NULL,
  [From] bigint NOT NULL,
  [To] bigint NOT NULL,
  [BranchId] int NOT NULL,
  [CheckInBefore] bigint NULL,
  [CheckOutAfter] bigint NULL,
  [TenantId] int NOT NULL,
  [CreatedBy] bigint NOT NULL,
  [ModifiedBy] bigint NULL,
  [ModifiedDate] datetime2 NULL
);


CREATE TABLE [TimeSheetShift] (
  [Id] bigint NOT NULL,
  [TimeSheetId] bigint NOT NULL,
  [ShiftIds] nvarchar(1000) NOT NULL,
  [RepeatDaysOfWeek] nvarchar(500) NULL,
  [DayOfWeek] ntext NOT NULL
);

CREATE TABLE [TimeSheet] (
  [Id] bigint NOT NULL,
  [EmployeeId] bigint NOT NULL,
  [StartDate] datetime2 NOT NULL,
  [EndDate] datetime2 NOT NULL,
  [IsRepeat] bit NULL,
  [RepeatType] tinyint NULL,
  [RepeatEachDay] tinyint NULL,
  [BranchId] int NOT NULL,
  [TenantId] int NOT NULL,
  [CreatedBy] bigint NOT NULL,
  [TimeSheetStatus] tinyint NOT NULL,
  [SaveOnDaysOffOfBranch] bit NOT NULL,
  [SaveOnHoliday] bit NOT NULL,
  [Note] nvarchar(250) NULL,
  [AutoGenerateClockingStatus] tinyint NOT NULL,
  [IsAppliedForNextWeeks] bit NULL,
  [ParentTimeSheetId] bigint NULL,
  [TimeSheetTrackingType] tinyint NULL
);


CREATE TABLE [Employee] (
  [Id] bigint NOT NULL,
  [NickName] nvarchar(255) NULL,
  [Name] nvarchar(255) NOT NULL,
  [DOB] datetime2 NULL,
  [IsActive] bit NOT NULL,
  [IdentityNumber] nvarchar(255) NULL,
  [MobilePhone] nvarchar(255) NULL,
  [UserId] bigint NULL,
  [DepartmentId] bigint NULL,
  [JobTitleId] bigint NULL,
  [IdentityKeyClocking] nvarchar(100) NULL,
  [AccountSecretKey] nvarchar(100) NULL,
  [TenantId] int NOT NULL,
  [BranchId] int NOT NULL,
  [CreatedBy] bigint NOT NULL,
  [ModifiedBy] bigint NULL,
  [StartWorkingDate] datetime2 NULL,
  [ModelDeviceClocking] nvarchar(100) NULL,
  [IsRecoverable] bit NOT NULL,
  [ReturnToWorkDate] datetime2 NULL
);

