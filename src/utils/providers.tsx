"use client";

import { api } from "./api";

interface PropsType {
  children: React.ReactNode;
}
type ProviderType = (props: PropsType) => JSX.Element;

function Providers({ children }: PropsType) {
  return <>{children}</>;
}

export default api.withTRPC(Providers) as ProviderType;
