import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import useAuth from "src/hooks/useAuth";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { SaveSettingsData, SaveSettingsSchema } from "src/validation/user";

const Settings = () => {
  const { auth, setAuth } = useAuth();
  const {
    formState: { errors },
    handleSubmit,
    register,
  } = useForm<SaveSettingsData>({
    defaultValues: {
      bmr: auth.user.bmr,
      caloriesPer100Steps: auth.user.caloriesPer100Steps,
    },
    resolver: zodResolver(SaveSettingsSchema),
  });
  const navigate = useNavigate();
  const axios = useAxiosPrivate();
  const createWeight = useMutation({
    mutationFn: async (data: SaveSettingsData) => {
      await axios.post("/user/settings", data);
      return data;
    },
    onSuccess: (data: SaveSettingsData) => {
      setAuth((prev) => ({
        ...prev,
        user: {
          ...prev.user,
          bmr: data.bmr,
          caloriesPer100Steps: data.caloriesPer100Steps,
        },
      }));
      navigate("/", { replace: true });
    },
    onError: (error: any) => {
      console.error("onError", error);
    },
  });

  const onSubmit = (data: SaveSettingsData) => {
    createWeight.mutate(data);
  };

  return (
    <section className='w-[555px] mx-auto'>
      <form id='foodForm' onSubmit={handleSubmit(onSubmit)} className='pt-3'>
        <div className='flex flex-col gap-2'>
          <section>
            <div className='label'>
              <span className='label-text'>Базовый обмен</span>
            </div>
            <input
              type='number'
              step='.1'
              className='input input-bordered w-full'
              {...register("bmr", {
                valueAsNumber: true,
              })}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.bmr?.message}
              </span>
            </div>
          </section>

          <section>
            <div className='label'>
              <span className='label-text'>Сжигаемые калории на 100 шагов</span>
            </div>
            <input
              type='number'
              step='.1'
              className='input input-bordered w-full'
              {...register("caloriesPer100Steps", {
                valueAsNumber: true,
              })}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.caloriesPer100Steps?.message}
              </span>
            </div>
          </section>

          <div className='flex justify-end gap-3 pt-5'>
            <button type='submit' className='btn btn-primary'>
              Сохранить
            </button>
          </div>
        </div>
      </form>
    </section>
  );
};

export default Settings;
