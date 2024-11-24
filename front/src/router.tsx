import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  Routes,
} from "react-router-dom";
import NotFound from "./routes/NotFound";
import Layout from "./routes/Layout";
import Home from "./routes/Home";
import Login from "./routes/Login";
import Register from "./routes/Register";
import Activities from "./routes/activities/Activities";
import AuthRequired from "./components/AuthRequired";
import PersistLogin from "./components/PersisLogin";
import StravaCallback from "./routes/StravaCallback";
import Foods from "./routes/foods/Foods";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route>
      <Route path='login' element={<Login />} />
      <Route path='register' element={<Register />} />

      <Route path='/' element={<Layout />}>
        <Route element={<PersistLogin />}>
          <Route element={<AuthRequired />}>
            <Route path='/' element={<Home />} />
            <Route path='/activities' element={<Activities />} />
            <Route path='/foods' element={<Foods />} />
            <Route path='/strava-oauth-callback' element={<StravaCallback />} />
            <Route path='*' element={<NotFound />} />
          </Route>
        </Route>
      </Route>
    </Route>
  )
);

export default router;
