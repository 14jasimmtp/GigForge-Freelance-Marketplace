package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Freelancer(
	api fiber.Router,
	profile *handler.ProfileHandler,
	project *handler.ProjectHandler,
	job *handler.JobsHandler,
) {
	profiles := api.Group("/profile")
	profiles.Get("", middlewares.AuthFreelancer, profile.GetProfile)
	profiles.Post("/education", middlewares.AuthFreelancer, profile.AddEducationDetails)
	profiles.Patch("/education", middlewares.AuthFreelancer, profile.UpdateEducation)
	profiles.Post("/experience", middlewares.AuthFreelancer, profile.AddExperience)
	profiles.Patch("/experience", middlewares.AuthFreelancer, profile.UpdateExperience)
	profiles.Delete("/experience", middlewares.AuthFreelancer, profile.RemoveExperience)
	profiles.Delete("/education", middlewares.AuthFreelancer, profile.DeleteEducation)
	profiles.Post("/description", middlewares.AuthFreelancer, profile.AddProfileDescription)
	profiles.Patch("/description", middlewares.AuthFreelancer, profile.EditProfileDescription)
	profiles.Put("/photo", middlewares.AuthFreelancer, profile.UpdateProfilePhoto)
	profiles.Post("/skill", middlewares.AuthFreelancer, profile.UpdateSkilltoProfile)

	jobs := api.Group("/job")
	jobs.Post("accept-offer/:offer_id", middlewares.AuthFreelancer, job.AcceptOffer)
	jobs.Post("send-proposal", middlewares.AuthFreelancer, job.SendProposal)
	jobs.Get("",job.GetJobs)
	jobs.Get("/:id",job.GetJob)
	jobs.Post("/invoice", middlewares.AuthFreelancer,job.SendInvoice)

}
