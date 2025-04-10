package paper

import (
	"fmt"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/schema/soo/wml"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"papergen/internal/models/question"
)

type Paper struct {
	gorm.Model
	Title       string         `gorm:"column:title;comment:'试卷标题'" json:"title"`
	Description string         `gorm:"column:description;comment:'试卷描述'" json:"description"`
	Questions   datatypes.JSON `gorm:"column:questions;type:json;comment:'试卷包含题目ID'" json:"questions"`
	Creator     string         `gorm:"column:creator;comment:'试卷创建人'" json:"creator"`
}

func GenerateDocxPaper(title string, questions []question.Question) error {
	doc := document.New()
	// 设置标题样式
	titlePara := doc.AddParagraph()
	titlePara.Properties().SetAlignment(wml.ST_JcCenter)
	run := titlePara.AddRun()
	run.Properties().SetSize(48)
	run.Properties().SetFontFamily("黑体")
	run.AddText(title)

	doc.AddParagraph() // 空行

	// 题型分组
	typeMap := map[string]string{
		"single_choice":   "一、单选题",
		"multiple_choice": "二、多选题",
		"true_false":      "三、判断题",
		"fill_blank":      "四、填空题",
		"short_answer":    "五、简答题",
	}

	grouped := map[string][]question.Question{}
	for _, q := range questions {
		grouped[q.QuestionType] = append(grouped[q.QuestionType], q)
	}

	// 遍历题型分类，逐个写入
	for _, typeKey := range []string{"single_choice", "multiple_choice", "true_false", "fill_blank", "short_answer"} {
		qs := grouped[typeKey]
		if len(qs) == 0 {
			continue
		}

		// 写入题型标题
		typeTitle := doc.AddParagraph()
		typeTitle.Properties().SetStyle("Heading2")
		typeTitle.AddRun().AddText(typeMap[typeKey])

		for i, q := range qs {
			// 写入题干
			p := doc.AddParagraph()
			qText := fmt.Sprintf("%d. %s", i+1, q.Question)
			p.AddRun().AddText(qText)

			// 如果是选择题，解析并写入选项
			if q.QuestionType == "single_choice" || q.QuestionType == "multiple_choice" {
				optP := doc.AddParagraph()
				optP.Properties().SetStyle("ListBullet")
				optP.AddRun().AddText(q.Options)
			}
		}

		doc.AddParagraph() // 空行
	}

	// 保存
	return doc.SaveToFile("tmp/" + title + ".docx")
}
