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
}

export default PermissionCheck;
