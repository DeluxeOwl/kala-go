interface PermissionCheck {
  subject: {
    type: string;
    name: string;
  };
  permission: string;
  resource: {
    type: string;
    name: string;
  };
  hasPermission?: boolean;
}

export default PermissionCheck;
