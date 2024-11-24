import { emptyUser } from "@/context/AuthProvider";
import useAuth from "@/hooks/useAuth";
import { Link } from "react-router-dom";

const NavBar = () => {
  const { auth, setAuth } = useAuth();

  return (
    <div className='navbar border-b border-gray-700 flex justify-between'>
      <div className='flex justify-start gap-10 items-end'>
        {/* <a href='/' id='#logo' className='text-primary text-4xl font-bold'>
          ActivityPeople
        </a> */}
        <Link to='/' id='#logo' className='text-primary text-4xl font-bold'>
          <i className='fa-solid fa-medal'></i>
        </Link>
        <nav className='text-primary text-xl font-bold'>
          <Link to='/foods'>Food</Link>
        </nav>
      </div>
      <ul className='menu menu-horizontal px-1 text-2xl flex gap-x-10'>
        <li>
          <details>
            <summary>{auth.user?.email}</summary>
            <ul className='p-2 bg-base-100 rounded-t-none z-10'>
              <li>
                <a href='/settings'>Settings</a>
              </li>
              <li onClick={() => setAuth(emptyUser)}>
                <a>Logout</a>
              </li>
            </ul>
          </details>
        </li>
      </ul>
    </div>
  );
};

export default NavBar;
