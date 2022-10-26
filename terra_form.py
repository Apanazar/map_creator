import matplotlib.pyplot as plt 
from random import choice


class RandomWalk():
	#Класс для генерирования случайных блужданий
	def __init__(self, num_points=5000):
		#Инициализация атрибутов блуждания
		self.num_points = num_points

		#Все блуждания начинаются с точки (0, 0)
		self.x_values = [0]
		self.y_values = [0]

	def fill_walk(self):
		#Вычисление точки блуждания
		#Шаги генерируются до достижения нужной длины
		while len(self.x_values) < self.num_points:
			#Определяю направление и длину перемещения
			x_direction = choice([1, -1])
			x_distance = choice([0, 1, 2, 3 , 4])
			x_step  = x_direction * x_distance

			y_direction = choice([1, -1])
			y_distance = choice([0, 1, 2, 3, 4])
			y_step = y_direction * y_distance

			#Отклонение нулевых перемещений
			if x_step == 0 and y_step == 0:
				continue

			#Вычисляю следующие значения х и у
			next_x = self.x_values[-1] + x_step
			next_y = self.y_values[-1] + y_step

			self.x_values.append(next_x)
			self.y_values.append(next_y)

while True:			
	rw = RandomWalk(300000)
	rw.fill_walk()

	point_numbers = list(range(rw.num_points))
	plt.scatter(rw.x_values, rw.y_values, s=10)
	plt.scatter(rw.x_values, rw.y_values, c=point_numbers, cmap=plt.cm.Blues, s=1)

	plt.scatter(0, 0, c='blue', edgecolor='black', s=100)
	plt.scatter(rw.x_values[-1], rw.y_values[-1], c='red', edgecolor='black', s=100)

	plt.savefig("mygraph.png")
	keep_running = input("Build a new wander? (Y/N): ")
	if keep_running == "n":
		break
