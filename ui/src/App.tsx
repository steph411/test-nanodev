import * as React from "react";
import { RouterProvider } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";
import { router } from "./router";
import { ReactQueryDevtools } from "react-query/devtools";
import AppLoader from "./components/AppLoader";

export const queryClient = new QueryClient();

const App = () => {
  return (
    <React.Suspense fallback={<AppLoader />}>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
        {/* <ReactQueryDevtools initialIsOpen={false} /> */}
      </QueryClientProvider>
    </React.Suspense>
  );
};

export default App;
