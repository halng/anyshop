/* eslint-disable @next/next/no-img-element */
import React, {
  forwardRef,
  useImperativeHandle,
  useRef,
} from "react";
import { AppTopbarRef } from "@/types";
import AppTopBarMenu from "./AppTopBarMenu";

const AppTopbar = forwardRef<AppTopbarRef>((props, ref) => {
  const menubuttonRef = useRef(null);
  const topbarmenuRef = useRef(null);
  const topbarmenubuttonRef = useRef(null);

  useImperativeHandle(ref, () => ({
    menubutton: menubuttonRef.current,
    topbarmenu: topbarmenuRef.current,
    topbarmenubutton: topbarmenubuttonRef.current,
  }));

  return (
    <div className="layout-topbar">
      <AppTopBarMenu />
    </div>
  );
});

AppTopbar.displayName = "AppTopbar";

export default AppTopbar;
