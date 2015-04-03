package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	// BuoyStation contains information for an individual station.
	MainDoc struct {
		ID               bson.ObjectId `bson:"_id,omitempty"`
		Class            string        `bson:"_class"`
		ReceivedOn       time.Time     `bson:"receivedOn"`
		ReceivedDocument Document      `bson:"document"`
	}

	Document struct {
		LocationGuid               string      `bson:"locationGuid"`
		LocationName               string      `bson:"locationName"`
		SurveyName                 string      `bson:"surveyName"`
		SurveyGuid                 string      `bson:"surveyGuid"`
		Question                   string      `bson:"question"`
		QuestionGuid               string      `bson:"questionGuid"`
		Answer                     []string    `bson:"answer"`
		AnswerGuid                 []string    `bson:"answerGuid"`
		Options                    []string    `bson:"options"`
		OptionGuids                []string    `bson:"optionGuids"`
		QuestionNumber             int         `bson:"questionNumber"`
		Skipped                    bool        `bson:"skipped"`
		SurveyQuestionType         string      `bson:"surveyQuestionType"`
		AccountId                  string      `bson:"accountId"`
		OriginatingTimestampString string      `bson:"originatingTimestampString"`
		DatePart                   DateParts   `bson:"dateParts"`
		SurveyResponseId           string      `bson:"surveyResponseId"`
		Validation                 SValidation `bson:"validation"`
	}

	DateParts struct {
		hourOfDayBase1 int    `bson:"hourOfDayBase1"`
		time           int    `bson:"time"`
		date           string `bson:"date"`
		dayOfWeek      string `bson:"dayOfWeek"`
	}

	SValidation struct {
		ValidationGuid  string             `bson:"validationGuid"`
		ValidationName  string             `bson:"validationName"`
		ValidationItems []SValidationItems `bson:"validationItems"`
	}

	SValidationItems struct {
		ValidationItemGuid string    `bson:"validationItemGuid"`
		ValidationItemName string    `bson:"validationItemName"`
		ValidationItemCode string    `bson:"validationItemCode"`
		Category           SCategory `bson:"category"`
	}

	SCategory struct {
		CategoryGuid string `bson:"categoryGuid"`
		CategoryName string `bson:"categoryName"`
	}

	//            "validationItems" : [
	//                {
	//                    "validationItemGuid" : "112be9b7-0902-45ac-abee-8eebee27c832",
	//                    "validationItemName" : "short Item2",
	//                    "validationItemCode" : "List Item2",
	//                    "category" : {
	//                        "categoryGuid" : "7460742b-ade7-43a9-adfd-7c2c65be5901",
	//                        "categoryName" : "Test Category"
	//                    }
	//                }
	//                {
	//                    "validationItemGuid" : "112be9b7-0902-45ac-abee-8eebee27c832",
	//                    "validationItemName" : "short Item2",
	//                    "validationItemCode" : "List Item2",
	//                    "category" : {
	//                        "categoryGuid" : "7460742b-ade7-43a9-adfd-7c2c65be5901",
	//                        "categoryName" : "Test Category"
	//                    }
	//                }

)

/*

    "_id" : ObjectId("53072a3ee4b039907a7ff4eb"),
    "_class" : "net.pager.lrs.analytics.aggregation.model.SurveyResponse",
    "receivedOn" : ISODate("2014-02-21T10:28:14.638Z"),
    "document" : {
        "locationGuid" : "7b5f4481-abbd-4d12-8320-0c2cfef2efc6",
        "locationName" : "A Pizza Ranch #567",
        "surveyName" : "All New Brand Survey",
        "surveyGuid" : "83b802b1-23c3-42de-b542-0455351e9ff3",
        "question" : "Delete free form",
        "questionGuid" : "5ddef910-8ea2-4830-8e58-543ce57ed011",
        "answer" : [
            "This need to be deleted 1"
        ],
        "answerGuid" : [],
        "options" : [],
        "optionGuids" : [],
        "questionNumber" : 0,
        "skipped" : false,
        "surveyQuestionType" : "LONGFORM",
        "accountId" : "7be7ced3-47b9-403a-b7f9-67bb19851f1a",
        "originatingTimestampString" : "2014-02-21T10:27:10.000+00:00",
        "dateParts" : {
            "hourOfDayBase1" : 10,
            "time" : 1027,
            "date" : 20140221,
            "dayOfWeek" : "Friday"
        },
        "surveyResponseId" : "99a8dcea-f9d8-2877-a1ed-6b8c4baa562e"
    }
}


*/
