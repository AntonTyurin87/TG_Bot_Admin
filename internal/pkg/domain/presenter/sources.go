package presenter

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/domain/menu"
	"TG_Bot_Admin/internal/pkg/domain/texts"
)

// CollbackKeyNameBySourceStep ...
func (p *presenter) CollbackKeyNameBySourceStep(step entity.Step) menu.KeyName {
	switch step {
	case entity.CreateSourceStep:
		return menu.KeyNameAddSourceNameRU
	case entity.SourceNameRuStep:
		return menu.KeyNameAddSourceNameENG
	case entity.SourceNameENGStep:
		return menu.KeyNameAddSourceAuthors
	case entity.SourceAuthorsRUStep:
		return menu.KeyNameAddSourceYear
	case entity.SourceYearStep:
		return menu.KeyNameAddSourceDescription
	case entity.SourceDescriptionStep:
		return menu.KeyNameAddSourceFile
	case entity.SourceLoadFileStep:
		return menu.KeyNameSourceSuccess
	default:
		return menu.UnknownKeyName
	}
}

// CollbackKeyNameBySourceType ...
func (p *presenter) CollbackKeyNameBySourceType(sourceType entity.SourceType) menu.KeyName {
	switch sourceType {
	case entity.BookSourceType:
		return menu.CreateBookSource
	case entity.ArticleSourceType:
		return menu.CreateArticleSource
	case entity.FragmentSourceType:
		return menu.CreateFragmentSource
	default:
		return menu.UnknownKeyName
	}
}

// InstructionsBySourceStep ...
func (p *presenter) InstructionsBySourceStep(step entity.Step) texts.Instructions {
	switch step {
	case entity.CreateSourceStep:
		return texts.InstructionsAddSourceNameRU
	case entity.SourceNameRuStep:
		return texts.InstructionsAddSourceNameENG
	case entity.SourceNameENGStep:
		return texts.InstructionsAddSourceAuthors
	case entity.SourceAuthorsRUStep:
		return texts.InstructionsAddSourceYear
	case entity.SourceYearStep:
		return texts.InstructionsAddSourceDescription
	case entity.SourceDescriptionStep:
		return texts.InstructionsAddSourceFile
	case entity.SourceLoadFileStep:
		return texts.InstructionsAddSourceSuccess
	default:
		return texts.InstructionsUnknown
	}
}
