package filler

import (
	model "github.com/Azamatttio/GoProject/Project/pkg/car/model"
)

func PopulateDatabase(models model.Models) error {
	for _, menu := range cars {
		models.Cars.Insert(&menu)
	}
	// TODO: Implement restaurants pupulation
	// TODO: Implement the relationship between restaurants and menus
	return nil
}

var cars = []model.Car{
	{Title: "BMW X5", Description: "Luxury SUV", Year: 2015},
	{Title: "Mercedes-Benz E-Class", Description: "Executive Sedan", Year: 2018},
	{Title: "Audi A4", Description: "Compact Executive Car", Year: 2016},
	{Title: "Toyota Camry", Description: "Mid-Size Sedan", Year: 2019},
	{Title: "Honda Civic", Description: "Compact Car", Year: 2020},
	{Title: "Ford F-150", Description: "Full-Size Pickup Truck", Year: 2017},
	{Title: "Chevrolet Silverado", Description: "Heavy-Duty Pickup Truck", Year: 2016},
	{Title: "Tesla Model S", Description: "Electric Luxury Sedan", Year: 2018},
	{Title: "Volkswagen Golf", Description: "Compact Hatchback", Year: 2017},
	{Title: "Subaru Outback", Description: "Crossover SUV", Year: 2022},
	{Title: "Lexus RX", Description: "Luxury Crossover SUV", Year: 2021},
	{Title: "Mazda CX-5", Description: "Compact Crossover SUV", Year: 2023},
	{Title: "Jeep Wrangler", Description: "Off-Road SUV", Year: 2015},
	{Title: "Hyundai Elantra", Description: "Economical Sedan", Year: 2014},
	{Title: "Kia Soul", Description: "Compact MPV", Year: 2013},
	{Title: "Nissan Altima", Description: "Mid-Size Sedan", Year: 2012},
	{Title: "Volvo XC90", Description: "Luxury Mid-Size SUV", Year: 2011},
	{Title: "Land Rover Range Rover", Description: "Luxury SUV", Year: 2010},
	{Title: "Porsche 911", Description: "Sports Car", Year: 2009},
	{Title: "Jaguar F-Pace", Description: "Luxury Compact SUV", Year: 2008},
	{Title: "Cadillac Escalade", Description: "Full-Size Luxury SUV", Year: 2007},
	{Title: "GMC Sierra", Description: "Full-Size Pickup Truck", Year: 2006},
	{Title: "Buick Encore", Description: "Subcompact Crossover SUV", Year: 2005},
	{Title: "Chrysler 300", Description: "Full-Size Sedan", Year: 2004},
	{Title: "Lincoln Navigator", Description: "Luxury Full-Size SUV", Year: 2003},
	{Title: "Infiniti Q50", Description: "Luxury Compact Sedan", Year: 2002},
	{Title: "Acura MDX", Description: "Luxury Mid-Size SUV", Year: 2001},
	{Title: "Ferrari 488", Description: "Exotic Sports Car", Year: 2000},
	{Title: "McLaren 720S", Description: "Super Car", Year: 1999},
	{Title: "Bentley Continental GT", Description: "Luxury Grand Tourer", Year: 1998},
	{Title: "Rolls-Royce Phantom", Description: "Ultra Luxury Sedan", Year: 1997},
}
