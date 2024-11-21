import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";
import NotFound from "./routes/NotFound";
import Layout from "./routes/Layout";
import Home from "./routes/Home";
import Login from "./routes/Login";
import Register from "./routes/Register";
import Activities from "./routes/Activities";
import AuthRequired from "./components/AuthRequired";
import PersistLogin from "./components/PersisLogin";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path='/' element={<Layout />}>
      <Route path='login' element={<Login />} />
      <Route path='register' element={<Register />} />

      <Route element={<PersistLogin />}>
        <Route element={<AuthRequired />}>
          <Route path='/' element={<Home />} />
          <Route path='/activities' element={<Activities />} />
          <Route path='*' element={<NotFound />} />
        </Route>
      </Route>
    </Route>
  )
);

export default router;
