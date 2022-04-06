package applicant

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"jobapp.com/m/common"
)

type ApplicantModel struct {
	Name           string                `json:"name" bson:"name" binding:"required"`
	Address        string                `json:"address" bson:"address" binding:"required"`
	PhoneNumber    string                `json:"phone_number" bson:"phone_number"`
	Email          string                `json:"email" bson:"email" binding:"required,email"`
	SocialMedia    []string              `json:"social_media" bson:"social_media"`
	AboutMe        string                `json:"about_me" bson:"about_me" binding:"required"`
	Skills         []string              `json:"skills" bson:"skills" binding:"required"`
	WorkExperience []WorkExperienceModel `json:"work_experience" bson:"work_experience" binding:"required"`
	Education      EducationModel        `json:"education" bson:"education" binding:"required"`
}

type WorkExperienceModel struct {
	OfficeName string   `json:"office_name" bson:"office_name" binding:"required"`
	Position   string   `json:"position" bson:"position" binding:"required"`
	JobDesks   []string `json:"job_desks" bson:"job_desks" binding:"required"`
}

type EducationModel struct {
	UniversityName string  `json:"university_name" bson:"university_name" binding:"required"`
	Degree         string  `json:"degree" bson:"degree" binding:"required"`
	Gpa            float64 `json:"gpa" bson:"gpa" binding:"required"`
}

func FindOneUser(key string, value string) (error, ApplicantModel) {

	var result bson.M
	var resultModel ApplicantModel

	err := common.ApplicantCollection.FindOne(
		context.TODO(),
		bson.D{{key, value}},
	).Decode(&result)
	if err != nil {
		return err, resultModel
	}
	bsonBytes, err := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &resultModel)
	if err != nil {
		return err, resultModel
	}
	return nil, resultModel
}

func InsertMongo(model ApplicantModel) error {
	_, err := common.ApplicantCollection.InsertOne(context.TODO(), model)
	return err
}

func DeleteMongo(name string) error {
	_, err := common.ApplicantCollection.DeleteOne(context.TODO(), bson.M{"name": name})
	return err
}

func UpdateMongo(model ApplicantModel) error {
	result, err := common.ApplicantCollection.UpdateOne(context.TODO(), bson.M{"email": model.Email}, bson.D{{Key: "$set", Value: model}})
	fmt.Println(result)
	return err
}
