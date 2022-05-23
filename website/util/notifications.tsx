import { showNotification } from "@mantine/notifications";
import { X } from "tabler-icons-react";

const showError = (errorMessage: string) => {
  showNotification({
    title: "Error",
    icon: <X size={18} />,
    color: "red",
    message: errorMessage,
    radius: "lg",
    style: { whiteSpace: "pre-line" },
  });
};

export { showError };
