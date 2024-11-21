import activityApi from "@/api/activitypeople";
import useAuth from "@/hooks/useAuth";
import { LoginData } from "@/validation/auth";
import { useForm } from "react-hook-form";
import { useLocation, useNavigate } from "react-router-dom";
import { userKey } from "..";

const Login = () => {
  const {
    formState: { errors },
    handleSubmit,
    register,
    setError,
  } = useForm<LoginData>();
  const { setAuth } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || "/";

  const onSubmit = async (c: LoginData) => {
    try {
      let res = await activityApi.login(c);
      const accessToken = res.data.accessToken;

      localStorage.setItem(userKey, JSON.stringify({ accessToken }));
      setAuth({ accessToken });

      navigate(from, { replace: true });
    } catch (e: any) {
      console.log("e error:", e);
      console.log("e code:", e.code);
    }
  };

  return (
    <div className='flex justify-center mt-[calc(100vh-100vh+8rem)]'>
      <div className='max-w-screen-sm w-full bg-base-300 py-10 px-16 rounded-xl'>
        <h1 className='text-center text-primary text-3xl mb-10'>Login</h1>

        <form
          hx-post='/login'
          className='space-y-5'
          onSubmit={handleSubmit(onSubmit)}
        >
          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Email</span>
            </div>
            <input
              type='email'
              placeholder='type your email'
              className='input input-bordered w-full'
              autoComplete='off'
              required
              {...register("email")}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.email?.message}
              </span>
            </div>
          </label>

          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Password</span>
            </div>
            <input
              type='password'
              placeholder='type your password'
              className='input input-bordered w-full'
              autoComplete='off'
              required
              {...register("password")}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.password?.message}
              </span>
            </div>
          </label>

          <button type='submit' className='btn btn-primary w-full'>
            login <i className='fa-solid fa-arrow-right'></i>
          </button>
          <div className='divider'>OR</div>
          <a href='/register' className='btn btn-outline w-full'>
            Register
          </a>
          {/* <a
            href='/login/provider/google'
            type='submit'
            className='btn btn-outline w-full'
          >
            Log in with Google <i className='fa-brands fa-google'></i>
          </a> */}
        </form>
      </div>
    </div>
  );
};

export default Login;
