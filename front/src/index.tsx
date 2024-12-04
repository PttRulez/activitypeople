import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import "./index.css";
import router from "./router";
import { AuthProvider } from "./context/AuthProvider";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

type localKey = string;
export const userKey: localKey = "acitivityuser";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

const queryClient = new QueryClient();

root.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <RouterProvider future={{ v7_startTransition: true }} router={router} />
      </AuthProvider>
    </QueryClientProvider>
  </React.StrictMode>
);
