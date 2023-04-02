import { FC, ReactNode, createContext, useEffect, useState } from "react";
import { Zmk } from "./layout";

export const ZmkContext = createContext<Zmk>({ layout: [], keys: {} });

interface ZmkProviderProps {
  children?: ReactNode;
}

export const ZmkProvider: FC<ZmkProviderProps> = ({ children }) => {
  const [zmk, setZmk] = useState<Zmk>({
    layout: [],
    keys: {},
  });

  useEffect(() => {
    fetch("http://localhost:5656/api/zmk")
      .then((r) => r.json())
      .then((j) => {
        console.log(j);
        setZmk(j);
      });
  }, []);

  return <ZmkContext.Provider value={zmk}>{children}</ZmkContext.Provider>;
};
