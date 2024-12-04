import useAuth from "src/hooks/useAuth";
import axios from "src/api/activitypeople";
import { RegisterData } from "src/validation/auth";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { UserInfo } from "src/context/AuthProvider";
import { AxiosResponse } from "axios";

const Register = () => {
  const {
    formState: { errors },
    handleSubmit,
    register,
    watch,
  } = useForm<RegisterData>();
  const { setAuth } = useAuth();
  const navigate = useNavigate();

  const onSubmit = async (c: RegisterData) => {
    try {
      let res: AxiosResponse<UserInfo> = await axios.post("/register", c);
      setAuth((prev) => res.data);

      navigate("/", { replace: true });
    } catch (e: any) {
      console.log("e error:", e);
      console.log("e code:", e.code);
    }
  };

  return (
    <div className='flex justify-center mt-[calc(100vh-100vh+8rem)]'>
      <div className='max-w-screen-sm w-full bg-base-300 py-10 px-16 rounded-xl'>
        <h1 className='text-center text-primary text-3xl mb-10'>Register</h1>

        <form
          hx-post='/login'
          className='space-y-5'
          onSubmit={handleSubmit(onSubmit)}
        >
          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Имя</span>
            </div>
            <input
              type='text'
              placeholder='ваше имя/ник на сайте'
              className='input input-bordered w-full'
              autoComplete='off'
              required
              {...register("name")}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.name?.message}
              </span>
            </div>
          </label>

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
              <span className='label-text'>Пароль</span>
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

          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Подвтерждение пароля</span>
            </div>
            <input
              type='password'
              placeholder='confirm your password'
              className='input input-bordered w-full'
              autoComplete='off'
              {...register("confirmPassword")}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.confirmPassword?.message}
              </span>
            </div>
          </label>

          <button type='submit' className='btn btn-primary w-full'>
            register <i className='fa-solid fa-arrow-right'></i>
          </button>
          <div className='divider'>OR</div>
          <a href='/login' className='btn btn-outline w-full'>
            Login
          </a>
        </form>
      </div>
    </div>
  );
};

export default Register;
