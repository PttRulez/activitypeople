import NavBar from "@/components/Navbar";
import { Outlet } from "react-router-dom";

export default function Layout() {
  return (
    <div className='container mx-auto'>
      <NavBar />
      <div className='py-12'>
        <Outlet />
      </div>
    </div>
  );
}
